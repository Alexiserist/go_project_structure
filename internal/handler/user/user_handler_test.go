package handler

import (
	"go_project_structure/internal/models/user"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)
var path = "/users"

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

func (m *MockUserService) UpdateUser(user models.User) (error){
	arg := m.Called(user)
	return arg.Error(0)
}

func TestGetAllUsers(t *testing.T){
	gin.SetMode(gin.TestMode)
	mockUserService := new(MockUserService)
	userHandler := NewUserHandler(mockUserService)
	router := gin.Default();
	router.GET(path, userHandler.GetUsers)
	mockUser := []models.User{
		{ID: 1, Username: "JohnDoe", Email: "johydefn@example.com", IsActive: true},
	}
	mockUserService.On("GetAllUsers").Return(mockUser, nil)
	req, _ := http.NewRequest(http.MethodGet,path, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "JohnDoe")
	mockUserService.AssertExpectations(t)
}

func TestGetUserByKey(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUserService := new(MockUserService)
	userHandler := NewUserHandler(mockUserService)

	router := gin.Default()
	router.GET(path + "/:id", userHandler.GetUserByKey)

	userID := uint(1)
	mockUser := models.User{ID: userID, Username: "JohnDoe", Email: "cornor@example.com"}

	mockUserService.On("FindOneByKey", userID).Return(mockUser, nil)

	req, _ := http.NewRequest(http.MethodGet, path + "/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "JohnDoe")
	assert.Contains(t, w.Body.String(), "cornor@example.com")

	mockUserService.AssertExpectations(t)
}

func TestCreateUser(t *testing.T){
	gin.SetMode(gin.TestMode)
	mockUserService := new(MockUserService)
	userHandler := NewUserHandler(mockUserService)

	router := gin.Default()
	router.POST(path,userHandler.CreateUser)
	mockBody := models.CreateUser{Username:"testing",Email: "johnydef@gmail",Password: "test",IsActive: true}
	mockUser := models.User{
		Username: mockBody.Username,
		Email:    mockBody.Email,
		Password: mockBody.Password,
		IsActive: mockBody.IsActive,
	}

	mockUserService.On("CreateUser", mock.Anything).Return(mockUser,nil)
    userPayload := `{"Username":"testing","Email":"johnydef@gmail","Password":"test","IsActive":true}`
    req, _ := http.NewRequest(http.MethodPost, path, strings.NewReader(userPayload))
    req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code);
	mockUserService.AssertCalled(t, "CreateUser", mock.Anything)
	mockUserService.AssertExpectations(t)

}

func TestUpdateUser(t *testing.T){
	gin.SetMode(gin.TestMode)
	mockUserService := new(MockUserService)
	UserHandler := NewUserHandler(mockUserService)

	router := gin.Default()
	router.PUT(path + "/:id", UserHandler.UpdateUser)
	mockUser := models.User{ID: uint(1), Username: "JohnDoe", Email: "john@example.com", Password: "test", IsActive: true}
	mockUserService.On("FindOneByKey", uint(1)).Return(mockUser, nil);
	mockUserService.On("UpdateUser",mockUser).Return(nil)

	userPayload := `{"Email":"john@example.com","Password":"test","IsActive":true}`
	req, _ := http.NewRequest(http.MethodPut, path+"/1", strings.NewReader(userPayload))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	mockUserService.AssertExpectations(t)
}


func TestDeleteUser(t *testing.T){
	gin.SetMode(gin.TestMode)
	mockUserService := new(MockUserService)
	UserHandler := NewUserHandler(mockUserService)

	router := gin.Default()
	router.DELETE(path + "/:id",UserHandler.DeleteUser)
	userID := uint(1)
	mockUser := models.User{ID: 1, Username: "JohnDoe", Email: "john@example.com"}
	mockUserService.On("FindOneByKey", userID).Return(mockUser, nil);
	mockUserService.On("DeleteUser", mock.AnythingOfType("models.User")).Return(nil)

	req, _ := http.NewRequest(http.MethodDelete, path + "/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w,req)

	assert.Equal(t, http.StatusOK, w.Code);
	mockUserService.AssertExpectations(t);
}