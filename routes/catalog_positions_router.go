package routes

import (
	"veterinary-employee/controllers"
	"veterinary-employee/db"
	"veterinary-employee/repositories"

	"github.com/labstack/echo/v4"
)

func BootstrapCatalogPositionsRoutes(v *echo.Group) {
	controller := &controllers.CatalogPositionsController{
		Repository: &repositories.CatalogPositionsRepository{
			Data: db.New(),
		},
	}

	v.GET("/catalog-positions", controller.GetAll)
	v.GET("/catalog-positions/:id", controller.GetById)
	v.POST("/catalog-positions", controller.Create)
	v.PATCH("/catalog-positions/:id", controller.UpdateById)
}
