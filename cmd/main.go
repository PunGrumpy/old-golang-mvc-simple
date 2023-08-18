package main

import (
	"github.com/PunGrumpy/golang-mvc/controller"
	"github.com/PunGrumpy/golang-mvc/pkg/env"
	"github.com/PunGrumpy/golang-mvc/service"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	soldierGroup := server.Group("/soldier")
	{
		soldierService := service.NewSoldierService()
		soldierController := controller.NewSoldierController(soldierService)

		soldierGroup.POST("/", soldierController.AddSoldier)
		soldierGroup.GET("/:id", soldierController.GetSoldierByID)
		soldierGroup.PUT("/:id", soldierController.UpdateSoldier)
	}
	port := env.GetEnvirontment("PORT")
	if port == "" {
		port = "8080"
	}
	if err := server.Run(":" + port); err != nil {
		panic(err.Error())
	}
}
