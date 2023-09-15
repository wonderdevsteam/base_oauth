package api

import (
	"github.com/go-chi/chi/v5"
)

func (transport *API) initRoutes(r chi.Router) {
	r.Get("/", transport.endpoints.Home)
	r.Get("/uptime", transport.endpoints.Uptime)

	r.NotFound(transport.endpoints.NotFound)
}