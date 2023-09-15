package main

import (
	"time"

	"go.uber.org/fx"
	uberzap "go.uber.org/zap"

	"oauth/core"
	"oauth/domain/usecases"
	"oauth/servers/api"
	"oauth/services/fxapp"
	"oauth/services/zap"
)

const timeout = 30 * time.Second

func main() {
	app := fx.New(
		fx.NopLogger, // remove for debug

		fx.Provide(
			func() *core.Config {
				return core.NewConfigFromYAML("conf/main.yml")
			},

			func(config *core.Config) *uberzap.Logger {
				log, _ := zap.NewZap(config.Logger)
				return log
			},

			usecases.New,
		),

		api.Transport,

		// fx.Invoke(validation.ValidateConfig),
		fx.Invoke(api.Run),
	)

	fxapp.Start(app, timeout)
	<-app.Done()
	fxapp.Shutdown(app, timeout)
}