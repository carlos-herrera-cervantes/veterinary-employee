package routes

import (
	"veterinary-employee/controllers"
	"veterinary-employee/db"
	"veterinary-employee/repositories"

	"github.com/labstack/echo/v4"
)

func BootstrapAvatarRoutes(v *echo.Group) {
	controller := &controllers.AvatarController{
		Repository: &repositories.AvatarRepository{
			Data: db.New(),
		},
	}

	v.GET("/employees/avatar/me", controller.GetMe)
	v.POST("/employees/avatar", controller.Create)
	v.PATCH("/employees/avatar/me", controller.UpdateMe)
	v.DELETE("/employees/avatar/me", controller.DeleteMe)
}
