package router

import "github.com/gin-gonic/gin"

func login(c *gin.Context) {

}
func logout(c *gin.Context) {}

func initAdminRouter(r *gin.Engine) {
	adminGroup := r.Group("/admin")

	{
		adminGroup.POST("/login", login)
		adminGroup.GET("/logout", logout)
	}
}
