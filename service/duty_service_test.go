package service

import (
	"testing"

	"github.com/PunGrumpy/golang-mvc/model"
	"github.com/stretchr/testify/assert"
)

const soldierNotFoundError = "Soldier Not Found"

func setupSoldierService() *soldierDutyService {
	return NewSoldierService().(*soldierDutyService)
}

func TestAddAndRetrieveSoldier(t *testing.T) {
	soldierService := setupSoldierService()

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

	t.Run("GetSoldierByID", func(t *testing.T) {
		soldier, err := soldierService.GetSoldierByID("1")
		assert.Nil(t, err)
		assert.Equal(t, newSoldier.Name, soldier.Name)
	})

	t.Run("GetErrorSoldierNotFound", func(t *testing.T) {
		soldier, err := soldierService.GetSoldierByID("2")
		assert.NotNil(t, err)
		assert.Nil(t, soldier)
		assert.Equal(t, soldierNotFoundError, err.Error())
	})
}

func TestUpdateSoldier(t *testing.T) {
	soldierService := setupSoldierService()

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

	updatedSoldier := &model.Soldier{
		ID:     1,
		Name:   "Bob",
		Rank:   "Sergeant",
		Wife:   "Eve",
		Salary: 40000,
		Home:   true,
		Car:    false,
	}

	t.Run("UpdateSoldier", func(t *testing.T) {
		err := soldierService.UpdateSoldier("1", updatedSoldier)
		assert.Nil(t, err)

		soldier, err := soldierService.GetSoldierByID("1")
		assert.Nil(t, err)
		assert.Equal(t, updatedSoldier.Name, soldier.Name)
	})

	t.Run("UpdateErrorSoldierNotFound", func(t *testing.T) {
		err := soldierService.UpdateSoldier("2", updatedSoldier)
		assert.NotNil(t, err)
		assert.Equal(t, soldierNotFoundError, err.Error())
	})
}
