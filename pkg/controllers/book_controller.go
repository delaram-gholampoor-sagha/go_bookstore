package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Delaram-Gholampoor-Sagha/go_bookstore/pkg/models"
	"github.com/Delaram-Gholampoor-Sagha/go_bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	// whatever we recieve from the database we want to turn it into json
	// its the json version of all the books that we found from the database
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	// helps us send somthing to the frontend ot ot postman
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Printf("err while parsing")
	}
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	// we recieve somthing from the user as a request now we want to parse that into somthing that our database will understand
	utils.ParseBody(r, CreateBook)
	// the same book that was sent to me by my user as a request
	b := CreateBook.CreateBook()
	// in the response i just converted that into json to be able to send it to the user
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeletetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Printf("err while parsing")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	// we are gonna take it in json format and pars in into (unmarshal it ) something that go understand ( our database understand )
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookid"]

	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Printf("err while parsing")
	}
	// we found the book by id
	bookDetails, db := models.GetBookById(ID)
	// now we are setting the new peorperies
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	// now we want to save these changes to database
	db.Save(&bookDetails)
	// now we know that it has to be json ...
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
