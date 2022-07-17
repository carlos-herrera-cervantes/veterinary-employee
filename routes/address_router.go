package routes

import (
	"veterinary-employee/controllers"
	"veterinary-employee/db"
	"veterinary-employee/repositories"

	"github.com/labstack/echo/v4"
)

func BootstrapAddressRoutes(v *echo.Group) {
	controller := &controllers.AddressController{
		Repository: &repositories.AddressRepository{
			Data: db.New(),
		},
	}

	v.GET("/employees/address/me", controller.GetMe)
	v.POST("/employees/address", controller.Create)
	v.PATCH("/employees/address/me", controller.UpdateMe)
}
