package httpfx

import (
	"net/http"

	"go.uber.org/fx"
)

var Module = fx.Option(
	fx.Provide(http.NewServeMux),
)
