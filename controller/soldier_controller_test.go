package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/PunGrumpy/golang-mvc-simple/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDutyService struct {
	mock.Mock
}

func (m *MockDutyService) AddSoldier(soldier *model.Soldier) error {
	args := m.Called(soldier)
	return args.Error(0)
}

func (m *MockDutyService) UpdateSoldier(soldierID string, updatedSoldier *model.Soldier) error {
	args := m.Called(soldierID, updatedSoldier)
	return args.Error(0)
}

func (m *MockDutyService) GetSoldierByID(soldierID string) (*model.Soldier, error) {
	args := m.Called(soldierID)
	return args.Get(0).(*model.Soldier), args.Error(1)
}

func TestAddSoldier(t *testing.T) {
	gin.SetMode(gin.TestMode)

	engine := gin.Default()
	mockService := new(MockDutyService)
	controller := NewSoldierController(mockService)
	engine.POST("/soldier", controller.AddSoldier)

	newSoldier := model.Soldier{
		Name:   "Alice",
		Rank:   "Sergeant",
		Salary: 40000,
	}

	payload, _ := json.Marshal(newSoldier)
	req, _ := http.NewRequest("POST", "/soldier", strings.NewReader(string(payload)))
	recorder := httptest.NewRecorder()

	mockService.On("AddSoldier", mock.AnythingOfType("*model.Soldier")).Return(nil)

	engine.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusCreated, recorder.Code)
	assert.Equal(t, "application/json; charset=utf-8", recorder.Header().Get("Content-Type"))

	var response struct {
		Soldier model.Soldier `json:"soldier"`
	}
	_ = json.Unmarshal(recorder.Body.Bytes(), &response)

	assert.Equal(t, newSoldier.Name, response.Soldier.Name)
}

func TestUpdateSoldier(t *testing.T) {
	gin.SetMode(gin.TestMode)

	engine := gin.Default()
	mockService := new(MockDutyService)
	controller := NewSoldierController(mockService)
	engine.PUT("/soldier/:id", controller.UpdateSoldier)

	updatedSoldier := model.Soldier{
		Name:   "Alice",
		Rank:   "Sergeant",
		Salary: 40000,
	}

	payload, _ := json.Marshal(updatedSoldier)
	req, _ := http.NewRequest("PUT", "/soldier/1", strings.NewReader(string(payload)))
	recorder := httptest.NewRecorder()

	mockService.On("UpdateSoldier", mock.AnythingOfType("string"), mock.AnythingOfType("*model.Soldier")).Return(nil)

	engine.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "application/json; charset=utf-8", recorder.Header().Get("Content-Type"))
}

func TestGetSoldierByID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	engine := gin.Default()
	mockService := new(MockDutyService)
	controller := NewSoldierController(mockService)
	engine.GET("/soldier/:id", controller.GetSoldierByID)

	newSoldier := model.Soldier{
		Name:   "Alice",
		Rank:   "Sergeant",
		Salary: 40000,
	}

	mockService.On("GetSoldierByID", mock.AnythingOfType("string")).Return(&newSoldier, nil)

	req, _ := http.NewRequest("GET", "/soldier/1", nil)
	recorder := httptest.NewRecorder()
	engine.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "application/json; charset=utf-8", recorder.Header().Get("Content-Type"))

	var response struct {
		Soldier model.Soldier `json:"soldier"`
	}
	_ = json.Unmarshal(recorder.Body.Bytes(), &response)

	assert.Equal(t, newSoldier.Name, response.Soldier.Name)
}
