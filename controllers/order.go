// controllers/order.go
package controllers

import (
	"encoding/json"
	"net/http"
	"restaurant-backend/database"
	"restaurant-backend/models"
)

func AddToCart(w http.ResponseWriter, r *http.Request) {
	var cartItem models.CartItem
	err := json.NewDecoder(r.Body).Decode(&cartItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Add cart item to the database or the current user's cart (for simplicity, just an example)
	collection := database.Client.Database("eldine").Collection("cart")
	_, err = collection.InsertOne(nil, cartItem)
	if err != nil {
		http.Error(w, "Failed to add to cart", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Item added to cart")
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Save order to database
	collection := database.Client.Database("eldine").Collection("orders")
	_, err = collection.InsertOne(nil, order)
	if err != nil {
		http.Error(w, "Failed to create order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Order created successfully")
}
