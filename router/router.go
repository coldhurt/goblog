package router

import (
	"github.com/coldhurt/goblog/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors())

	initAdminRouter(r)
	initArticleRouter(r)

	return r
}
