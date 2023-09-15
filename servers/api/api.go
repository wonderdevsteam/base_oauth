package api

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"

	"oauth/domain/usecases"
	"oauth/servers"
	"oauth/servers/api/config"
	"oauth/servers/api/routes"
	"oauth/servers/cors"
)

type API struct {
	httpConfig *config.Config
	endpoints  *routes.Endpoints
	httpServer *http.Server
}

func NewHTTP(httpConfig *config.Config, logger *zap.Logger, uc *usecases.Usecases) servers.TCP {
	return &API{
		httpConfig: httpConfig,
		endpoints:  routes.NewEndpoints(logger, uc),
	}
}

func (transport *API) ListenAndServe() error {
	r := chi.NewRouter()
	transport.initMiddlewares(r)
	transport.initRoutes(r)

	transport.httpServer = &http.Server{
		Addr:    transport.httpConfig.GetAddress(),
		Handler: r,
	}

	return transport.httpServer.ListenAndServe()
}

func (transport *API) Shutdown() error {
	return transport.httpServer.Shutdown(context.Background())
}

func (transport *API) GetAddress() string {
	return transport.httpConfig.GetAddress()
}

func (transport *API) initMiddlewares(router chi.Router) {
	router.Use(middleware.StripSlashes)
	router.Use(middleware.SetHeader("X-Frame-Options", "deny"))
	router.Use(middleware.Timeout(5 * time.Minute)) //nolint:gomnd
	router.Use(middleware.Heartbeat("/ping"))
	router.Use(middleware.RealIP) // Add this middleware ONLY if the service is behind a proxy or gateway !

	cors.InitGoCHI(transport.httpConfig.CORS, router)

	if transport.httpConfig.NoCache {
		router.Use(middleware.NoCache)
	}

	router.Use(middleware.AllowContentType(
		"application/json",
	))

}