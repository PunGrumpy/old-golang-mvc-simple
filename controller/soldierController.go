package controller

import (
	"strconv"

	"github.com/PunGrumpy/golang-mvc/service"
	"github.com/gin-gonic/gin"
)

type SoldierController interface {
	Eat(c *gin.Context)
}

type soldierController struct {
	dutyService service.DutyService
}

func SoldierHandler(dutyService *service.DutyService) SoldierController {
	return &soldierController{
		dutyService: *dutyService,
	}
}

func (s *soldierController) Eat(c *gin.Context) {
	commision, err := strconv.Atoi(c.Param("commision"))
	if err != nil {
		c.AbortWithStatus(400)
	}
	s.dutyService.EatTax(c, commision)
}
