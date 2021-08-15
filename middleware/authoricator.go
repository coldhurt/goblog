package middleware

import (
	"net/http"

	"github.com/coldhurt/goblog/models"
	"github.com/coldhurt/goblog/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Authenticator(c *gin.Context) {
	token := c.GetHeader("Authorization")
	service := service.JWTAuthService()
	t, err := service.ValidateToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.Message{Msg: "token is expired"})
	}

	if _, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		// c.JSON(http.StatusOK, models.Message{Msg: "ok", Data: claims})
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, models.Message{Msg: "token is expired"})
	}
}
