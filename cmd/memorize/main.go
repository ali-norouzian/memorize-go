package main

import (
	"memorize/internal/router"
	"memorize/pkg/swagger"

	"go.uber.org/fx"
)

func main() {
	swagger.InitSwagger()

	fx.New(
		router.Module,
	).Run()
}
