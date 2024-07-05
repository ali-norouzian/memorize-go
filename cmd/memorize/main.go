package main

import (
	"memorize/internal/router"
	"memorize/pkg/swagger"

	"go.uber.org/fx"
)

// @title Memorize API
// @version 1.0
//// @description This is a sample server for Swagger documentation.
//// @termsOfService http://swagger.io/terms/

//// @contact.name API Support
//// @contact.url http://www.swagger.io/support
//// @contact.email support@swagger.io

//// @license.name Apache 2.0
//// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath

// @externalDocs.description Open Swagger UI
// @externalDocs.url http://localhost:8000/swagger/index.html

func main() {
	swagger.InitSwagger()

	fx.New(
		router.Module,
	).Run()
}
