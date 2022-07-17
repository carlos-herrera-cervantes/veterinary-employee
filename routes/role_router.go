package routes

import (
	"veterinary-employee/controllers"
	"veterinary-employee/db"
	"veterinary-employee/repositories"

	"github.com/labstack/echo/v4"
)

func BootstrapRoleRoutes(v *echo.Group) {
	controller := &controllers.RoleController{
		Repository: &repositories.RoleRepository{
			Data: db.New(),
		},
	}

	v.GET("/employees/roles", controller.GetAll)
	v.POST("/employees/roles", controller.Create)
	v.PATCH("/employees/roles/:id", controller.Update)
}
