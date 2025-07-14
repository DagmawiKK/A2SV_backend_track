package controllers

import (
	"fmt"
	"library_management/models"
	"library_management/services"
)

type LibraryController struct {
	Library *services.Library
}

func NewLibraryController(library *services.Library) *LibraryController {
	return &LibraryController{
		Library: library,
	}
}

func (lc *LibraryController) AddBookHandler(id int, title, author string) {
	book := models.Book{
		ID:     id,
		Title:  title,
		Author: author,
		Status: "Available",
	}
	lc.Library.AddBook(book)
	fmt.Println("Book added successfully.")
}

func (lc *LibraryController) RemoveBookHandler(bookID int) {
	lc.Library.RemoveBook(bookID)
	fmt.Println("Book removed successfully.")
}

func (lc *LibraryController) BorrowBookHandler(bookID, memberID int) {
	err := lc.Library.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book borrowed successfully.")
	}
}

func (lc *LibraryController) ReturnBookHandler(bookID, memberID int) {
	err := lc.Library.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book returned successfully.")
	}
}

func (lc *LibraryController) ListAvailableBooksHandler() {
	availableBooks := lc.Library.ListAvailableBooks()
	fmt.Println("Available Books:")
	for _, book := range availableBooks {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}

func (lc *LibraryController) ListBorrowedBooksHandler(memberID int) {
	borrowedBooks := lc.Library.ListBorrowedBooks(memberID)
	fmt.Println("Borrowed Books:")
	for _, book := range borrowedBooks {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}
