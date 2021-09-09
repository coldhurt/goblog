package service

import (
	"fmt"
	"log"

	"github.com/coldhurt/goblog/db"
	"github.com/coldhurt/goblog/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllArticles(page int64, pageSize int64) []models.Article {
	var articles []models.Article

	client, collection, ctx, cancel := db.GetCollection("article")
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

func GetArticleById(id string) (models.Article, error) {
	var article models.Article
	client, collection, ctx, cancel := db.GetCollection("article")
	defer cancel()
	defer client.Disconnect(ctx)
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return article, err
	}
	result := collection.FindOne(ctx, bson.M{"_id": objID})
	result.Decode(&article)
	return article, nil
}

func CreateArticle(title string, content string) (primitive.ObjectID, error) {
	_, collection, ctx, _ := db.GetCollection("article")

	article := models.Article{Title: title, Content: content}
	article.ID = primitive.NewObjectID()

	_, err := collection.InsertOne(ctx, article)
	if err != nil {
		log.Printf("create article error %v", err)
		return article.ID, err
	}

	return article.ID, nil
}

func UpdateArticle(id, title, content string) error {
	if id == "" || title == "" || content == "" {
		return fmt.Errorf("bad request")
	}
	_, collection, ctx, _ := db.GetCollection("article")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid id")
	}
	_, err = collection.UpdateOne(ctx,
		bson.M{"_id": objID},
		bson.D{{
			"$set", bson.M{
				"title":   title,
				"content": content,
			}}})

	if err != nil {
		log.Printf("update article error %v", err)
		return err
	}

	return nil
}

func DeleteArticle(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid id")
	}
	_, collection, ctx, _ := db.GetCollection("article")
	_, err = collection.DeleteOne(ctx,
		bson.M{"_id": objID},
	)
	if err != nil {
		return fmt.Errorf("update article error %s", err.Error())
	}
	return nil
}
