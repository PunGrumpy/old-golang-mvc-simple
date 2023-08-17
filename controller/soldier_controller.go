package controller

import (
	"net/http"
	"strconv"

	"github.com/PunGrumpy/golang-mvc/model"
	"github.com/PunGrumpy/golang-mvc/service"
	"github.com/gin-gonic/gin"
)

type SoldierController interface {
	AddSoldier(c *gin.Context)
	UpdateSoldier(c *gin.Context)
	GetSoldierByID(c *gin.Context)
}

type soldierController struct {
	dutyService service.DutyService
}

func NewSoldierController(dutyservice service.DutyService) SoldierController {
	return &soldierController{
		dutyService: dutyservice,
	}
}

func (s *soldierController) AddSoldier(c *gin.Context) {
	var newSoldier model.Soldier
	if err := c.ShouldBindJSON(&newSoldier); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Soldier Input"})
		return
	}

	id := strconv.Itoa(newSoldier.ID)
	if _, found := s.dutyService.GetSoldierByID(id); found == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Soldier ID Already Exists"})
		return
	}

	s.dutyService.AddSoldier(&newSoldier)
	c.JSON(http.StatusCreated, gin.H{"soldier": newSoldier})
}

func (s *soldierController) UpdateSoldier(c *gin.Context) {
	id := c.Param("id")
	var updateSoldier model.Soldier
	if err := c.ShouldBindJSON(&updateSoldier); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Soldier Input"})
		return
	}

	if err := s.dutyService.UpdateSoldier(id, &updateSoldier); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func (s *soldierController) GetSoldierByID(c *gin.Context) {
	id := c.Param("id")
	soldierInfo, err := s.dutyService.GetSoldierByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"soldier": soldierInfo})
}
