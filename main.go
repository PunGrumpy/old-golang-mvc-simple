package main

import (
	"github.com/PunGrumpy/golang-mvc/controller"
	"github.com/PunGrumpy/golang-mvc/service"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	soldier := server.Group("/soldier")
	{
		soldier.POST("/eat/:commision", func(c *gin.Context) {
			var dutyService service.DutyService = service.SoldierDutyService(c)
			var soldierController controller.SoldierController = controller.SoldierHandler(&dutyService)
			soldierController.Eat(c)
		})
	}
	port := "8080"
	server.Run(":" + port)
}
