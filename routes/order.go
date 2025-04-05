// routes/order.go
package routes

import (
	"net/http"
	"restaurant-backend/controllers"
)

func OrderRoutes() {
	http.HandleFunc("/add_to_cart", controllers.AddToCart)
	http.HandleFunc("/create_order", controllers.CreateOrder)
}
