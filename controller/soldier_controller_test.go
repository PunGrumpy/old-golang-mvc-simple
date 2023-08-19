package controller

import (
	"encoding/json"
	"errors"
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

func ServerMock(t *testing.T) (*gin.Engine, *MockDutyService) {
	gin.SetMode(gin.TestMode)

	server := gin.Default()
	mockService := new(MockDutyService)
	mockController := NewSoldierController(mockService)

	serverGroup := server.Group("/soldier")
	{
		serverGroup.POST("", mockController.AddSoldier)
		serverGroup.PUT("/:id", mockController.UpdateSoldier)
		serverGroup.GET("/:id", mockController.GetSoldierByID)
	}

	return server, mockService
}

func TestAddSoldier(t *testing.T) {
	engine, mockService := ServerMock(t)

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

func TestAddSoldierWithInvalidPayload(t *testing.T) {
	engine, _ := ServerMock(t)

	req, _ := http.NewRequest("POST", "/soldier", strings.NewReader("invalid payload"))
	recorder := httptest.NewRecorder()

	engine.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.Equal(t, "application/json; charset=utf-8", recorder.Header().Get("Content-Type"))
}

func TestAddSoldierWithAlreadyExistSoldier(t *testing.T) {
	engine, mockService := ServerMock(t)

	newSoldier := model.Soldier{
		Name:   "Alice",
		Rank:   "Sergeant",
		Salary: 40000,
	}

	payload, _ := json.Marshal(newSoldier)
	req, _ := http.NewRequest("POST", "/soldier", strings.NewReader(string(payload)))
	recorder := httptest.NewRecorder()

	mockService.On("AddSoldier", mock.AnythingOfType("*model.Soldier")).Return(errors.New("Soldier already exists"))

	engine.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.Equal(t, "application/json; charset=utf-8", recorder.Header().Get("Content-Type"))
}

func TestUpdateSoldier(t *testing.T) {
	engine, mockService := ServerMock(t)

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

func TestUpdateSoldierWithInvalidPayload(t *testing.T) {
	engine, _ := ServerMock(t)

	req, _ := http.NewRequest("PUT", "/soldier/1", strings.NewReader("invalid payload"))
	recorder := httptest.NewRecorder()

	engine.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.Equal(t, "application/json; charset=utf-8", recorder.Header().Get("Content-Type"))
}

func TestUpdateSoldierWithNotFoundSoldier(t *testing.T) {
	engine, mockService := ServerMock(t)

	updatedSoldier := model.Soldier{
		Name:   "Alice",
		Rank:   "Sergeant",
		Salary: 40000,
	}

	payload, _ := json.Marshal(updatedSoldier)
	req, _ := http.NewRequest("PUT", "/soldier/1", strings.NewReader(string(payload)))
	recorder := httptest.NewRecorder()

	mockService.On("UpdateSoldier", mock.AnythingOfType("string"), mock.AnythingOfType("*model.Soldier")).Return(errors.New("Soldier not found"))

	engine.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.Equal(t, "application/json; charset=utf-8", recorder.Header().Get("Content-Type"))
}

func TestGetSoldierByID(t *testing.T) {
	engine, mockService := ServerMock(t)

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

func TestGetSoldierByIDWithNotFoundSoldier(t *testing.T) {
	engine, mockService := ServerMock(t)

	newSoldier := model.Soldier{
		Name:   "Alice",
		Rank:   "Sergeant",
		Salary: 40000,
	}

	mockService.On("GetSoldierByID", mock.AnythingOfType("string")).Return(&newSoldier, errors.New("Soldier not found"))

	req, _ := http.NewRequest("GET", "/soldier/1", nil)
	recorder := httptest.NewRecorder()
	engine.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusNotFound, recorder.Code)
	assert.Equal(t, "application/json; charset=utf-8", recorder.Header().Get("Content-Type"))
}
