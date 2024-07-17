package handler

import (
	// "errors"
	"go_project_structure/internal/models"
	"go_project_structure/internal/services"
	"go_project_structure/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// @Summary Get all users
// @Description Get all users
// @Tags Users
// @Accept  json
// @Produce  json
// @Success 200 {object} object{status=int,message=string,data=[]models.User}
// @Router /users [get]
// @Security BearerAuth
func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		utils.RespondWithStatusMessage(c, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	utils.ResponseWithStatusNessageData(c, http.StatusOK, http.StatusText(http.StatusOK), users)
}

// @Summary Get Users By Key
// @Description Get Users By Key
// @Tags Users
// @Accept  json
// @Param id path int true "User ID"
// @Produce  json
// @Success 200 {object} object{status=int,message=string}
// @Router /users/{id} [get]
// @Security BearerAuth
func (h *UserHandler) GetUserByKey(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	users, err := h.userService.FindOneByKey(uint(id))
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			utils.RespondWithStatusMessage(c, http.StatusNotFound, "No data found")
		} else {
			utils.RespondWithStatusMessage(c, http.StatusInternalServerError, err.Error())
		}
		return
	}
	utils.ResponseWithStatusNessageData(c, http.StatusOK, http.StatusText(http.StatusOK), users)
}

// @Summary Create users
// @Description Create users
// @Tags Users
// @Accept  json
// @Param body body models.CreateUser true "Request Body"
// @Produce  json
// @Success 200 {object} object{status=int,message=string,data=models.User}
// @Router /users [post]
// @Security BearerAuth
func (h *UserHandler) CreateUser(c *gin.Context) {
	var createUser models.CreateUser
	if err := c.ShouldBindBodyWithJSON(&createUser); err != nil {
		utils.RespondWithStatusMessage(c, http.StatusBadRequest, "Invalid request payload")
		return
	}
	users := models.User{
		Username: createUser.Username,
		Email:    createUser.Email,
		Password: createUser.Password,
		IsActive: createUser.IsActive,
	}
	users, err := h.userService.CreateUser(users)
	if err != nil {
		utils.RespondWithStatusMessage(c, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	utils.ResponseWithStatusNessageData(c, http.StatusOK, http.StatusText(http.StatusOK), users)
}

// @Summary Delete Users By Key
// @Description Delete Users By Key
// @Tags Users
// @Accept  json
// @Param id path int true "User ID"
// @Produce  json
// @Success 200 {object} object{status=int,message=string}
// @Router /users/{id} [delete]
// @Security BearerAuth
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := h.userService.FindOneByKey(uint(id))
	log.Println(user)
	if err != nil {
		utils.RespondWithStatusMessage(c, http.StatusNotFound, http.StatusText(http.StatusNotFound))
	}
	err = h.userService.DeleteUser(user)
	if err != nil {
		utils.RespondWithStatusMessage(c, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	utils.RespondWithStatusMessage(c, http.StatusOK, http.StatusText(http.StatusOK))
}
