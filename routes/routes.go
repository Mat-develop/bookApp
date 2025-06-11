package routes

import (
	"book-register-app/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router) {
	bookController := controllers.BookController{}

	router.HandleFunc("/books", bookController.CreateBook).Methods("POST")
	router.HandleFunc("/books", bookController.GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", bookController.GetBook).Methods("GET")
	router.HandleFunc("/books/{id}", bookController.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", bookController.DeleteBook).Methods("DELETE")
}
