package service

import (
	"testing"

	"github.com/PunGrumpy/golang-mvc/model"
	"github.com/stretchr/testify/assert"
)

func TestAddSoldier(t *testing.T) {
	soldierService := NewSoldierService()

	newSoldier := &model.Soldier{
		ID:     1,
		Name:   "Alice",
		Rank:   "Sergeant",
		Wife:   "Eve",
		Salary: 40000,
		Home:   true,
		Car:    false,
	}

	soldierService.AddSoldier(newSoldier)

	soldier, err := soldierService.GetSoldierByID("1")
	assert.Nil(t, err)
	assert.Equal(t, newSoldier.Name, soldier.Name)
}

func TestUpdateSoldier(t *testing.T) {
	soldierDutyService := NewSoldierService()

	newSoldier := &model.Soldier{
		ID:     1,
		Name:   "Alice",
		Rank:   "Sergeant",
		Wife:   "Eve",
		Salary: 40000,
		Home:   true,
		Car:    false,
	}

	soldierDutyService.AddSoldier(newSoldier)

	updatedSoldier := &model.Soldier{
		ID:     1,
		Name:   "Bob",
		Rank:   "Sergeant",
		Wife:   "Eve",
		Salary: 40000,
		Home:   true,
		Car:    false,
	}

	err := soldierDutyService.UpdateSoldier("1", updatedSoldier)
	assert.Nil(t, err)
}

func TestUpdateErrorSoldierNotFound(t *testing.T) {
	soldierDutyService := NewSoldierService()

	updatedSoldier := &model.Soldier{
		ID:     1,
		Name:   "Bob",
		Rank:   "Sergeant",
		Wife:   "Eve",
		Salary: 40000,
		Home:   true,
		Car:    false,
	}

	err := soldierDutyService.UpdateSoldier("2", updatedSoldier)
	assert.NotNil(t, err)
	assert.Equal(t, "Soldier Not Found", err.Error())
}

func TestGetSoldierByID(t *testing.T) {
	soldierDutyService := NewSoldierService()

	newSoldier := &model.Soldier{
		ID:     1,
		Name:   "Alice",
		Rank:   "Sergeant",
		Wife:   "Eve",
		Salary: 40000,
		Home:   true,
		Car:    false,
	}

	soldierDutyService.AddSoldier(newSoldier)

	soldier, err := soldierDutyService.GetSoldierByID("1")
	assert.Nil(t, err)
	assert.Equal(t, newSoldier.Name, soldier.Name)
}

func TestGetErrorSoldierNotFound(t *testing.T) {
	soldierDutyService := NewSoldierService()

	newSoldier := &model.Soldier{
		ID:     1,
		Name:   "Alice",
		Rank:   "Sergeant",
		Wife:   "Eve",
		Salary: 40000,
		Home:   true,
		Car:    false,
	}

	soldierDutyService.AddSoldier(newSoldier)

	soldier, err := soldierDutyService.GetSoldierByID("2")
	assert.NotNil(t, err)
	assert.Nil(t, soldier)
	assert.Equal(t, "Soldier Not Found", err.Error())
}
