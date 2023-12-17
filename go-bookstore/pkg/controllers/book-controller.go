package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Naveenchand06/go-projects/go-bookstore/pkg/models"
	"github.com/Naveenchand06/go-projects/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID, err := strconv.ParseInt(vars["bookId"], 0, 0)
	if err != nil {
		fmt.Println("Error while parsing ", err)
	}
	bookDetails, _ := models.GetBookyById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	newBook := models.Book{}
	utils.ParseBody(r, newBook)
	b := newBook.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId, err := strconv.ParseInt(vars["bookId"], 0, 0)
	if err != nil {
		fmt.Println("Error while parsing error")

	}
	book := models.DeleteBook(bookId)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func UpdateBook(w http.ResponseWriter, r *http.Request) {
	updateBook := models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId, err := strconv.ParseInt(vars["bookId"], 0, 0)
	if err != nil {
		fmt.Println("Error whilr parsing")
	}
	bookDetails, db := models.GetBookyById(bookId)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}