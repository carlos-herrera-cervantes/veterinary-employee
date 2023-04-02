package server

import (
	"fmt"

	"veterinary-employee/routes"
	"veterinary-employee/settings"

	"github.com/labstack/echo/v4"
)

func BootstrapServer() {
	e := echo.New()
	v1 := e.Group(fmt.Sprintf("%s%s", settings.InitializeApp().BasePath, "/v1"))

	routes.BootstrapProfileRoutes(v1)
	routes.BootstrapRoleRoutes(v1)
	routes.BootstrapAddressRoutes(v1)
	routes.BootstrapAvatarRoutes(v1)

	if err := e.Start(fmt.Sprintf(":%d", settings.InitializeApp().ServerPort)); err != nil {
		panic(err.Error())
	}
}
