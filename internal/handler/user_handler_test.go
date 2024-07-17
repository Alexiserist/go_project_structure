package handler

import (
	"go_project_structure/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) GetAllUsers() ([]models.User, error){
	args := m.Called()
	return args.Get(0).([]models.User), args.Error(1)
}

func (m *MockUserService) FindOneByKey(id uint) (models.User, error){
	args := m.Called(id)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *MockUserService) CreateUser(user models.User) (models.User, error){
	args := m.Called(user)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *MockUserService) DeleteUser(user models.User) (error){
	args := m.Called(user)
	return args.Error(0)
}

func TestGetAllUsers(t *testing.T){
	gin.SetMode(gin.TestMode)
	mockUserService := new(MockUserService)
	userHandler := NewUserHandler(mockUserService)
	router := gin.Default();
	router.GET("/users", userHandler.GetUsers)
	mockUser := []models.User{
		{ID: 1, Username: "JohnDoe", Email: "john@example.com", IsActive: true},
		{ID: 2, Username: "JaneDoe", Email: "JaneDoe@example.com", IsActive: true},
	}
	mockUserService.On("GetAllUsers").Return(mockUser, nil)
	req, _ := http.NewRequest(http.MethodGet, "/users", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "JohnDoe")
	assert.Contains(t, w.Body.String(), "JaneDoe")
	mockUserService.AssertExpectations(t)
}

func TestGetUserByKey(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUserService := new(MockUserService)
	userHandler := NewUserHandler(mockUserService)

	router := gin.Default()
	router.GET("/users/:id", userHandler.GetUserByKey)

	userID := uint(1)
	mockUser := models.User{ID: userID, Username: "JohnDoe", Email: "john@example.com"}

	mockUserService.On("FindOneByKey", userID).Return(mockUser, nil)

	req, _ := http.NewRequest(http.MethodGet, "/users/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "JohnDoe")
	assert.Contains(t, w.Body.String(), "john@example.com")

	mockUserService.AssertExpectations(t)
}
