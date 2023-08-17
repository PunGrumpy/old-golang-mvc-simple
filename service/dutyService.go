package service

import (
	"net/http"

	"github.com/PunGrumpy/golang-mvc/cores"
	"github.com/gin-gonic/gin"
)

type DutyService interface {
	EatTax(c *gin.Context, commision int)
}

type soldier cores.Soldier

func SoldierDutyService(c *gin.Context) DutyService {
	var soilder cores.Soldier
	if err := c.ShouldBind(&soilder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		c.Abort()
		return nil
	}
	return &soldier{
		Rank:       soilder.Rank,
		Wife:       soilder.Wife,
		Salary:     soilder.Salary,
		Home:       soilder.Home,
		Car:        soilder.Car,
		Corruption: soilder.Corruption,
	}
}

func (s *soldier) EatTax(c *gin.Context, commision int) {
	if s.Rank != "General" {
		s.Salary -= commision
	}

	if s.Corruption {
		s.promote(s, "Elite")
	}

	c.JSON(http.StatusOK, gin.H{
		"rank":       s.Rank,
		"wife":       s.Wife,
		"salary":     s.Salary,
		"home":       s.Home,
		"car":        s.Car,
		"corruption": s.Corruption,
	})
}

func (s *soldier) promote(soldier *soldier, newRank string) {
	soldier.Rank = newRank
	soldier.Salary *= 2
	soldier.Car = true
	soldier.Home = true
	soldier.Wife = "Beautiful"
}
