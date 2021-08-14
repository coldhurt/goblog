package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	initAdminRouter(r)
	initArticleRouter(r)

	return r
}
