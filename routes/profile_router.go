package routes

import (
	"veterinary-employee/controllers"
	"veterinary-employee/db"
	"veterinary-employee/repositories"

	"github.com/labstack/echo/v4"
)

func BootstrapProfileRoutes(v *echo.Group) {
	controller := &controllers.ProfileController{
		Repository: &repositories.ProfileRepository{
			Data: db.New(),
		},
	}

	v.GET("/employees/profiles", controller.GetAll)
	v.GET("/employees/profiles/:id", controller.GetById)
	v.GET("/employees/profiles/me", controller.GetMe)
	v.PATCH("/employees/profiles/me", controller.UpdateMe)
}
