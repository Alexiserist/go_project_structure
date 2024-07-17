package auth

import (
	"go_project_structure/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService AuthService
}

func NewAuthHandler(authService AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// @Summary Authorization Login
// @Description Login
// @Tags Auth
// @Accept  json
// @Param body body auth.UserLogin true "Request Body"
// @Produce  json
// @Success 200 {object} auth.AccessTokenResponse
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var userLogin UserLogin;
	if err := c.ShouldBindBodyWithJSON(&userLogin); err != nil {
		utils.RespondWithStatusMessage(c, http.StatusBadRequest, "Invalid request payload")
        return
    }
	token, err := h.authService.LoginHandler(userLogin.Username,userLogin.Password);
	if err != nil {
		switch err {
		case utils.ErrInvalidCredentials:
			utils.RespondWithStatusMessage(c, http.StatusUnauthorized, err.Error())
		case utils.ErrTokenGeneration:
			utils.RespondWithStatusMessage(c, http.StatusInternalServerError, err.Error())
		default:
			utils.RespondWithStatusMessage(c, http.StatusInternalServerError, "Internal server error")
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"message":  "OK",
		"data": 	token,
	})
}