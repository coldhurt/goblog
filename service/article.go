package service

import "github.com/coldhurt/goblog/models"

var articleList = []models.Article{
	{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

func GetAllArticles() []models.Article {
	return articleList
}

func GetArticleById(id int) models.Article {
	var article models.Article
	for _, item := range articleList {
		if item.ID == id {
			article = item
			break
		}
	}
	return article
}
