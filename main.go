package main

import (
	"echogo/database"
	_ "echogo/docs"
	"echogo/midlewares"
	"echogo/rute"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// @title Zigmades Keuangan
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost.swagger.io
// @BasePath /
func main() {
	db := database.Connect()
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.RequestID())
	e.Validator = &midlewares.CustomValidator{Validator: validator.New()}
	e.Use(midlewares.ContextDB(db))
	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	apiGroup := e.Group("/api")
	rute.Ruteprivate(apiGroup)
	umumGrup := e.Group("/api")
	rute.RuteUmum(umumGrup)

	e.Logger.Fatal(e.Start(":8000"))
}
