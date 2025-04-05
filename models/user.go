// models/user.go
package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username string             `json:"username" bson:"username"`
	Password string             `json:"password" bson:"password"`
	Role     string             `json:"role" bson:"role"` // "user" or "rider"
}

type LoginCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
