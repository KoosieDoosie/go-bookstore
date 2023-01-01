package routes

import (
	"github.com/KoosieDoosie/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")           //create function in controllers dir. Create new book
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")               //create function in controllers dir. Get books
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")   //Get books by ID
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")    //UPDATE books by ID
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE") //Get books by ID

}
