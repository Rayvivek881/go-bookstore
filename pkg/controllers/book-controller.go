package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rayvivek881/go-bookstore/pkg/models"
	"github.com/rayvivek881/go-bookstore/pkg/utils"
	"net/http"
	"strconv"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBook := models.GetAllBooks()
	res, _ := json.Marshal(newBook)
	w.Header().Set("Content-Type", "application/json") // set the content type to json
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // get all params
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		fmt.Println("error while parsing bookId", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	createBook := &models.Book{}   // create a new instance of Book by interface
	utils.ParseBody(r, createBook) // parse the request body to the interface
	b := createBook.CreateBook()
	res, _ := json.Marshal(b) // convert the response to json
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		fmt.Println("error while parsing bookId", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	bookDetails := models.DeleteBook(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook) // parse the request body to the interface
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		fmt.Println("error while parsing bookId", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	bookDetails, db := models.GetBookById(ID)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails) // save the updated book details
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
