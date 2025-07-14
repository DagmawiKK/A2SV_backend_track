package main

import (
	"library_management/controllers"
	"library_management/models"
	"library_management/services"
)

func main() {
	library := services.NewLibrary()
	controller := controllers.NewLibraryController(library)

	controller.AddBookHandler(1, "The Go Programming Language", "Alan")
	controller.AddBookHandler(2, "Clean Code", "Robert")

	library.Members[1] = models.Member{
		ID:   1,
		Name: "Abebe Chala",
	}

	controller.ListAvailableBooksHandler()

	controller.BorrowBookHandler(1, 1)

	controller.ListBorrowedBooksHandler(1)

	controller.ReturnBookHandler(1, 1)

	controller.ListAvailableBooksHandler()
}
