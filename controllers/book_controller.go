package controllers

import (
	"book-register-app/database"
	"book-register-app/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type BookController struct{}

func (bc *BookController) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := database.Connect()
	defer db.Close()

	_, err = db.Exec("INSERT INTO books (title, author, published_year, image_url) VALUES (?, ?, ?, ?)", book.Title, book.Author, book.PublishedYear, book.ImageURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func (bc *BookController) GetBooks(w http.ResponseWriter, r *http.Request) {
	db := database.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT id, title, author, published_year, image_url FROM books")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.PublishedYear, &book.ImageURL); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		books = append(books, book)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func (bc *BookController) GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	db := database.Connect()
	defer db.Close()

	var book models.Book
	err := db.QueryRow("SELECT id, title, author, published_year, image_url FROM books WHERE id = ?", id).Scan(&book.ID, &book.Title, &book.Author, &book.PublishedYear, &book.ImageURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func (bc *BookController) UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := database.Connect()
	defer db.Close()

	_, err = db.Exec("UPDATE books SET title = ?, author = ?, published_year = ?, image_url = ? WHERE id = ?", book.Title, book.Author, book.PublishedYear, book.ImageURL, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (bc *BookController) DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	db := database.Connect()
	defer db.Close()

	_, err := db.Exec("DELETE FROM books WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
