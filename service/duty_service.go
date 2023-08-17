package service

import (
	"errors"
	"strconv"
	"sync"

	"github.com/PunGrumpy/golang-mvc/model"
)

var (
	soldiersMap   = make(map[string]*model.Soldier)
	soldiersMutex sync.RWMutex
)

type DutyService interface {
	AddSoldier(soldier *model.Soldier)
	UpdateSoldier(soldierID string, updatedSoldier *model.Soldier) error
	GetSoldierByID(soliderID string) (*model.Soldier, error)
}

type soldierDutyService struct {
	soldiers map[string]*model.Soldier
}

func NewSoldierService() DutyService {
	return &soldierDutyService{
		soldiers: soldiersMap,
	}
}

func (s *soldierDutyService) AddSoldier(soldier *model.Soldier) {
	soldiersMutex.Lock()
	defer soldiersMutex.Unlock()

	s.soldiers[strconv.Itoa(soldier.ID)] = soldier
}

func (s *soldierDutyService) UpdateSoldier(soldierID string, updatedSoldier *model.Soldier) error {
	soldiersMutex.Lock()
	defer soldiersMutex.Unlock()

	if _, found := s.soldiers[soldierID]; !found {
		return errors.New("Soldier Not Found")
	}

	s.soldiers[soldierID] = updatedSoldier
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
