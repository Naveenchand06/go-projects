package routes

import (
	"github.com/Naveenchand06/go-projects/go-mongo-api/controllers"
	"github.com/gorilla/mux"
)

func RegisterUserRoutes(router *mux.Router) {
	// * Create User Route 
	router.HandleFunc("/user", controllers.CreateUserController).Methods("POST")

	// * Get All User Route 
	router.HandleFunc("/users", controllers.GetAllUserController).Methods("GET")

	// * Get User by ID Route 
	router.HandleFunc("/users/{id}", controllers.GetUserByIDController).Methods("GET")

	// * Update User By ID Route 
	router.HandleFunc("/users/{id}", controllers.UpdateUserByIdController).Methods("PUT")

	// * Delete User Route 
	router.HandleFunc("/users/{id}", controllers.DeleteUserByIdController).Methods("DELETE")
}
