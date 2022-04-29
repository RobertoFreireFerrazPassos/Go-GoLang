package main

// go mod init singleresponsability

import (
	"fmt"
	"solid/singleresponsability/books"
)

func main() {
	bookStore := books.BookStore{}
	bookStore.AddBook("Book 1 - Author 1")
	fmt.Println(bookStore.AddBook("Book 2 - Author 2"))
	fmt.Println(bookStore)
}
