package router

import (
	"net/http"

	"github.com/coldhurt/goblog/models"
	"github.com/coldhurt/goblog/service"
	"github.com/gin-gonic/gin"
)

type bodyLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func login(c *gin.Context) {
	var d bodyLogin
	err := c.ShouldBindJSON(&d)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{Msg: "bad request"})
		return
	}
	token, err := service.Login(d.Username, d.Password)
	if err != nil {
		c.JSON(http.StatusOK, models.Message{Msg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.Message{Msg: "ok", Data: map[string]string{"token": token}})
}

func logout(c *gin.Context) {

}

func initAdminRouter(r *gin.Engine) {
	adminGroup := r.Group("/admin")
	{
		adminGroup.POST("/login", login)
		adminGroup.GET("/logout", logout)
	}
}
