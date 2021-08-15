package service

import (
	"context"
	"log"

	"github.com/coldhurt/goblog/db"
	"github.com/coldhurt/goblog/models"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getCollection() (*mongo.Client, *mongo.Collection, context.Context, context.CancelFunc) {
	client, ctx, cancel := db.GetConnection()
	collection := client.Database(viper.GetString("MONGODB_DATABASE")).Collection("article")
	return client, collection, ctx, cancel
}

func GetAllArticles(page int64, pageSize int64) []models.Article {
	var articles []models.Article

	client, collection, ctx, cancel := getCollection()
	defer cancel()
	defer client.Disconnect(ctx)

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * pageSize)
	findOptions.SetLimit(pageSize)
	cur, err := collection.Find(ctx, bson.D{}, findOptions)
	if err != nil {
		log.Printf("Article find err %v", err)
	}
	for cur.Next(ctx) {

		// create a value into which the single document can be decoded
		var elem models.Article
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		articles = append(articles, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(ctx)
	return articles
}

func GetArticleById(id int) models.Article {
	var article models.Article
	return article
}

func CreateArticle(title string, content string) (primitive.ObjectID, error) {
	client, collection, ctx, cancel := getCollection()
	defer cancel()
	defer client.Disconnect(ctx)

	article := models.Article{Title: title, Content: content}
	article.ID = primitive.NewObjectID()

	_, err := collection.InsertOne(ctx, article)
	if err != nil {
		log.Printf("create article error %v", err)
		return article.ID, err
	}

	return article.ID, nil
}
