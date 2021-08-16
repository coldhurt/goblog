package middleware

import (
	"fmt"
	"net/http"

	"github.com/coldhurt/goblog/models"
	"github.com/coldhurt/goblog/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Authenticator() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		service := service.JWTAuthService()
		t, err := service.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, models.Message{Msg: "token is expired"})
		}

		if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
			fmt.Printf("request username %s\n", claims["username"])
			c.Set("user", claims)
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.Message{Msg: "token is expired"})
		}
	}
}
