package service

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/coldhurt/goblog/db"
	"github.com/coldhurt/goblog/models"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
)

type JWTService interface {
	GenerateToken(email string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type authCustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
	issure    string
}

//auth-jwt
func JWTAuthService() JWTService {
	return &jwtServices{
		secretKey: getSecretKey(),
		issure:    "Bikash",
	}
}

func getSecretKey() string {
	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func (service *jwtServices) GenerateToken(email string) string {
	claims := &authCustomClaims{
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    service.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("invalid token %s", token.Header["alg"])

		}
		return []byte(service.secretKey), nil
	})
}

func Login(username string, password string) (string, error) {
	service := JWTAuthService()
	var admin *models.Admin

	client, ctx, cancel := db.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	collection := client.Database(viper.GetString("MONGODB_DATABASE")).Collection("admin")
	result := collection.FindOne(ctx, bson.M{"username": username})
	if result == nil {
		return "", errors.New("no this user")
	}
	err := result.Decode(&admin)

	if err != nil {
		log.Printf("Failed marshalling %v", err)
		return "", errors.New("no this user")
	}
	log.Printf("Find admin: %v", admin)
	return service.GenerateToken(admin.Username), nil
}
