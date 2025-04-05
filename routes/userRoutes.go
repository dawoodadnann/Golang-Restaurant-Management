// routes/user.go
package routes

import (
	"net/http"
	"restaurant-backend/controllers"
)

func UserRoutes() {
	http.HandleFunc("/signup", controllers.SignUp)
	http.HandleFunc("/login", controllers.Login)
}
