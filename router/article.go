package router

import "github.com/gin-gonic/gin"

func createArticle(c *gin.Context) {}
func listArticle(c *gin.Context)   {}
func editArticle(c *gin.Context)   {}
func removeArticle(c *gin.Context) {}

func initArticleRouter(r *gin.Engine) {
	articleGroup := r.Group("/article")

	{
		articleGroup.POST("/create", createArticle)
		articleGroup.PUT("/edit", editArticle)
		articleGroup.GET("/list", listArticle)
		articleGroup.DELETE("/remove", removeArticle)
	}
}
