package service

import (
	"testing"

	"github.com/PunGrumpy/golang-mvc-simple/model"
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

	if err := soldierService.AddSoldier(newSoldier); err != nil {
		t.Error(err)
	}

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

	t.Run("AlreadyExistsError", func(t *testing.T) {
		err := soldierService.AddSoldier(newSoldier)
		assert.NotNil(t, err)
		assert.Equal(t, "Soldier ID Already Exists", err.Error())
	})

	soldiersMap = make(map[string]*model.Soldier)
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

	if err := soldierService.AddSoldier(newSoldier); err != nil {
		t.Error(err)
	}

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

	soldiersMap = make(map[string]*model.Soldier)
}

func TestDeleteSoldier(t *testing.T) {
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

	if err := soldierService.AddSoldier(newSoldier); err != nil {
		t.Error(err)
	}

	t.Run("DeleteSoldier", func(t *testing.T) {
		err := soldierService.DeleteSoldierByID("1")
		assert.Nil(t, err)

		soldier, err := soldierService.GetSoldierByID("1")
		assert.NotNil(t, err)
		assert.Nil(t, soldier)
		assert.Equal(t, soldierNotFoundError, err.Error())
	})

	t.Run("DeleteErrorSoldierNotFound", func(t *testing.T) {
		err := soldierService.DeleteSoldierByID("2")
		assert.NotNil(t, err)
		assert.Equal(t, soldierNotFoundError, err.Error())
	})

	soldiersMap = make(map[string]*model.Soldier)
}
