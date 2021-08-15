package router

import (
	"net/http"

	"github.com/coldhurt/goblog/middleware"
	"github.com/coldhurt/goblog/models"
	"github.com/coldhurt/goblog/service"
	"github.com/gin-gonic/gin"
)

type ListArticleBody struct {
	PageNo   int64 `json:"page_no"`
	PageSize int64 `json:"page_size"`
}

type CreateArticleBody struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func createArticle(c *gin.Context) {
	var body CreateArticleBody
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{Msg: "bad request"})
		return
	}
	id, err := service.CreateArticle(body.Title, body.Content)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{Msg: "bad request"})
		return
	}
	c.JSON(http.StatusOK, models.Message{Msg: "ok", Data: id})
}

func listArticle(c *gin.Context) {
	var body ListArticleBody
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{Msg: "bad request"})
		return
	}
	articles := service.GetAllArticles(body.PageNo, body.PageSize)
	c.JSON(http.StatusOK, models.Message{Msg: "ok", Data: articles})
}
func editArticle(c *gin.Context)   {}
func removeArticle(c *gin.Context) {}

func initArticleRouter(r *gin.Engine) {
	articleGroup := r.Group("/article")

	{
		articleGroup.POST("/create", middleware.Authenticator, createArticle)
		articleGroup.PUT("/edit", middleware.Authenticator, editArticle)
		articleGroup.POST("/list", listArticle)
		articleGroup.DELETE("/remove", middleware.Authenticator, removeArticle)
	}
}
