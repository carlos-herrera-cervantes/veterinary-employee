package server

import (
	"fmt"

	"veterinary-employee/settings"

	"github.com/labstack/echo/v4"
)

func BootstrapServer() {
	e := echo.New()
	//v1 := e.Group("api/v1")

	e.Start(fmt.Sprintf(":%d", settings.InitializeApp().ServerPort))
}
