package models

type Admin struct {
	ID       string `json:"id", bson:"_id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
