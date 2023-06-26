package routes

import (
	"veterinary-employee/controllers"
	"veterinary-employee/db"
	"veterinary-employee/repositories"
	"veterinary-employee/services"
	"veterinary-employee/singleton"

	"github.com/labstack/echo/v4"
)

func BootstrapProfileRoutes(v *echo.Group) {
	controller := &controllers.ProfileController{
		Repository: &repositories.ProfileRepository{
			Data: db.New(),
		},
		KafkaService: &services.KafkaService{
			Producer: singleton.NewProducer(),
		},
	}

	v.GET("/profiles", controller.GetAll)
	v.GET("/profiles/:id", controller.GetById)
	v.GET("/profiles/me", controller.GetMe)
	v.PATCH("/profiles/me", controller.UpdateMe)
	v.PATCH("/profiles/:id", controller.UpdateById)
}
