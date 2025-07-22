package main

import (
	"fmt"
	"log"
	"net/http"

	httpAdapter "github.com/sky0621/tips-go/library_selection/mock/mockery/adapter/http"
	"github.com/sky0621/tips-go/library_selection/mock/mockery/infrastructure/repository/memory"
	"github.com/sky0621/tips-go/library_selection/mock/mockery/usecase"
)

func main() {
	// Create repository
	userRepo := memory.NewUserRepository()

	// Create use case
	userUseCase := usecase.NewUserUseCase(userRepo)

	// Create HTTP handler
	userHandler := httpAdapter.NewUserHandler(userUseCase)

	// Register routes
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			userHandler.GetAllUsers(w, r)
		case http.MethodPost:
			userHandler.CreateUser(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			userHandler.GetUser(w, r)
		case http.MethodPut:
			userHandler.UpdateUser(w, r)
		case http.MethodDelete:
			userHandler.DeleteUser(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Start server
	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
