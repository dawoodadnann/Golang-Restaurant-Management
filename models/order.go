// models/order.go
package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CartItem struct {
	ProductID primitive.ObjectID `json:"product_id" bson:"product_id"`
	Quantity  int                `json:"quantity" bson:"quantity"`
}

type Order struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id"`
	Items     []CartItem         `json:"items" bson:"items"`
	TotalCost float64            `json:"total_cost" bson:"total_cost"`
	Status    string             `json:"status" bson:"status"` // "pending", "completed"
}
