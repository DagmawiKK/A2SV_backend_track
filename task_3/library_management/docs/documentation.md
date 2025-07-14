# Library Management System

## Overview
This is a simple console-based LMS written in Go. It allows users to manage books, borrow and return them, and list available and borrowed books.

## Features
- **AddBook**: Add a new book to the library.
- **RemoveBook**: Remove an existing book from the library.
- **BorrowBook**: Borrow a book if it's available.
- **ReturnBook**: Return a borrowed book.
- **ListAvailableBooks**: List all books that are currently available for borrowing.
- **ListBorrowedBooks**: List all books borrowed by a specific member.

## System Components
- **Book**: Represents a book in the library.
- **Member**: Represents a library member.
- **Library**: Manages the collection of books and members.
- **LibraryManager**: Interface for performing operations on the library.
