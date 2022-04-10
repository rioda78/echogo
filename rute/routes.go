package rute

import (
	"echogo/kontroller"
	"echogo/midlewares"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
)

func Ruteprivate(e *echo.Group) {
	e.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, this is echo2!")
	})

	/*	e.GET("/pegawai", kontroller.FetchAllPegawai, midlewares.IsAuthenticated)
		e.POST("/pegawai", kontroller.StorePegawai, midlewares.IsAuthenticated)
		e.PUT("/pegawai", kontroller.UpdatePegawai, midlewares.IsAuthenticated)
		e.DELETE("/pegawai", kontroller.DeletePegawai, midlewares.IsAuthenticated)

		e.GET("/generate-hash/:password", kontroller.GenerateHashPassword)
		e.POST("/login", kontroller.CheckLogin)

		e.GET("/test-struct-validation", kontroller.TestStructValidation)
		e.GET("/test-variable-validation", kontroller.TestVariableValidation)*/
}

func RuteUmum(e *echo.Group) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "selamat datang di home!")
	})
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.POST("/register", kontroller.Register)
	e.POST("/login", kontroller.Login)
	e.POST("/logout", kontroller.Logout, midlewares.Otorisasi)
}
