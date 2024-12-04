package middleware

import (
	"go_project_structure/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)


type AuthMiddleware interface {
	AuthorizationMiddleware(c *gin.Context)
}

type authMiddleware struct{
	authService auth.AuthService
}

func NewAuthMiddleware() AuthMiddleware {
	return &authMiddleware{
		authService: auth.NewAuthRepository(),
	}
}

func (m *authMiddleware) AuthorizationMiddleware(c *gin.Context) {
	header := c.Request.Header.Get("Authorization");
	if header == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Authorization header missing",
		})
		c.Abort()
		return
	}

	token := strings.TrimPrefix(header, "Bearer ")
	if err := m.authService.ValidateToken(token); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Invalid Token",
		})
		c.Abort()
		return
	}
	c.Next()
}
