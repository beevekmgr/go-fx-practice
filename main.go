package main

import (
	"go.uber.org/fx"

	"github.com/go-fx-practice/bundlefx"
	"github.com/go-fx-practice/httphandler"
)

func main() {

	fx.New(
		bundlefx.Module,
		fx.Invoke(httphandler.New),
	).Run()

}
