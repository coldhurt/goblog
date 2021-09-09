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
	"golang.org/x/crypto/bcrypt"
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

var s JWTService

//auth-jwt
func JWTAuthService() JWTService {
	return &jwtServices{
		secretKey: getSecretKey(),
		issure:    "Bikash",
	}
}

func getJwtService() JWTService {
	if s == nil {
		s = JWTAuthService()
	}
	return s
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

func Login(username, password string) (string, error) {
	service := JWTAuthService()
	var admin *models.Admin

	client, ctx, _ := db.GetConnection()
	// defer cancel()
	// defer client.Disconnect(ctx)
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

	log.Printf("Find admin: %v", admin.Username)

	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password)); err != nil {
		return "", fmt.Errorf("invalid username or password")
	}
	return service.GenerateToken(admin.Username), nil
}

func UpdatePassword(username, oldPassword, newPassword string) error {
	if username == "" || oldPassword == "" || len(newPassword) < 8 {
		return fmt.Errorf("bad request")
	}
	var admin models.Admin
	_, collection, ctx, _ := db.GetCollection("admin")
	result := collection.FindOne(ctx, bson.M{"username": username})
	err := result.Decode(&admin)

	if err != nil {
		log.Printf("Failed marshalling %v", err)
		return fmt.Errorf("no this user %s", username)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(oldPassword)); err != nil {
		return fmt.Errorf("incorrect password")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)

	if err != nil {
		log.Printf("%v", err)
		return errors.New("hash error")
	}

	fmt.Printf("update password for user id %s\n", admin.ID)

	updateRes, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": admin.ID},
		bson.D{
			{"$set", bson.M{"password": string(hashedPassword)}},
		},
	)

	if err != nil {
		return err
	}

	fmt.Printf("Updated %v Documents!\n", updateRes.ModifiedCount)

	return nil
}
