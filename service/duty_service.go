package service

import (
	"errors"
	"strconv"
	"sync"

	"github.com/PunGrumpy/golang-mvc-simple/model"
)

var (
	soldiersMap   = make(map[string]*model.Soldier)
	soldiersMutex sync.RWMutex
)

type DutyService interface {
	AddSoldier(soldier *model.Soldier) error
	UpdateSoldier(soldierID string, updatedSoldier *model.Soldier) error
	GetSoldierByID(soliderID string) (*model.Soldier, error)
	DeleteSoldierByID(soldierID string) error
}

type soldierDutyService struct {
	soldiers map[string]*model.Soldier
}

func NewSoldierService() DutyService {
	return &soldierDutyService{
		soldiers: soldiersMap,
	}
}

func (s *soldierDutyService) AddSoldier(soldier *model.Soldier) error {
	soldiersMutex.Lock()
	defer soldiersMutex.Unlock()

	if _, found := s.soldiers[strconv.Itoa(soldier.ID)]; found {
		return errors.New("Soldier ID Already Exists")
	}

	s.soldiers[strconv.Itoa(soldier.ID)] = soldier
	return nil
}

func (s *soldierDutyService) UpdateSoldier(soldierID string, updatedSoldier *model.Soldier) error {
	soldiersMutex.Lock()
	defer soldiersMutex.Unlock()

	if _, found := s.soldiers[soldierID]; !found {
		return errors.New("Soldier Not Found")
	}

	soldier := s.soldiers[soldierID]

	if updatedSoldier.Name != "" {
		soldier.Name = updatedSoldier.Name
	}
	if updatedSoldier.Rank != "" {
		soldier.Rank = updatedSoldier.Rank
	}
	if updatedSoldier.Salary != 0 {
		soldier.Salary = updatedSoldier.Salary
	}
	soldier.Home = updatedSoldier.Home
	soldier.Car = updatedSoldier.Car
	soldier.Corruption = updatedSoldier.Corruption

	return nil
}

func (s *soldierDutyService) GetSoldierByID(soldierID string) (*model.Soldier, error) {
	soldiersMutex.RLock()
	defer soldiersMutex.RUnlock()

	soldier, found := s.soldiers[soldierID]
	if !found {
		return nil, errors.New("Soldier Not Found")
	}
	return soldier, nil
}

func (s *soldierDutyService) DeleteSoldierByID(soldierID string) error {
	soldiersMutex.Lock()
	defer soldiersMutex.Unlock()

	if _, found := s.soldiers[soldierID]; !found {
		return errors.New("Soldier Not Found")
	}

	delete(s.soldiers, soldierID)
	return nil
}
