package main

import (
	"github.com/PunGrumpy/golang-mvc/controller"
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
	port := "8080"
	server.Run(":" + port)
}
