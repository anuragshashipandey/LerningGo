package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/anurag/bookstore/pkg/models"
	"github.com/anurag/bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBook()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	w.Header().Set("Content-Type", "pkglication/json")
	if err != nil {
		fmt.Println("error while parsing")
		w.WriteHeader(http.StatusNotFound)
	} else {
		bookDetails, _ := models.GetBookById(Id)
		res, _ := json.Marshal(bookDetails)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	w.Header().Set("Content-Type", "pkglication/json")
	if err != nil {
		fmt.Println("error while parsing")
		w.WriteHeader(http.StatusNotFound)
	} else {
		book := models.DeleteBook(Id)
		w.WriteHeader(http.StatusOK)
		res, _ := json.Marshal(book)
		w.Write(res)
	}
}

func UpdateBookById(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
		w.WriteHeader(http.StatusNotFound)
	} else {
		book, db := models.GetBookById(Id)
		if updateBook.Name != "" {
			book.Name = updateBook.Name
		}
		if updateBook.Author != "" {
			book.Author = updateBook.Author
		}
		if updateBook.Publication != "" {
			book.Publication = updateBook.Author
		}
		db.Save(&book)
		res, _ := json.Marshal(book)
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}

}
