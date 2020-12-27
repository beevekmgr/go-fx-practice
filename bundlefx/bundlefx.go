package bundlefx

import (
	"context"
	"net/http"

	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/go-fx-practice/configfx"
	"github.com/go-fx-practice/httpfx"
	"github.com/go-fx-practice/loggerfx"
)

func registerHooks(
	lifecycle fx.Lifecycle,
	logger *zap.SugaredLogger, cfg *configfx.Config, mux *http.ServeMux,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				logger.Info("Listening on ", cfg.ApplicationConfig.Address)
				go http.ListenAndServe(cfg.ApplicationConfig.Address, mux)
				return nil
			},
			OnStop: func(context.Context) error {
				return logger.Sync()
			},
		},
	)
}

var Module = fx.Options(
	configfx.Module,
	loggerfx.Module,
	httpfx.Module,
	fx.Invoke(registerHooks),
)
