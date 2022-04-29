package main

// go mod init singleresponsability

import (
	"fmt"
	"mainmodule/books"
)

func main() {
	bookStore := books.BookStore{}
	book1 := books.Book{Title: "Book 1 - Author 1", Category: books.Adventure}
	book2 := books.Book{Title: "Book 2 - Author 2", Category: books.Biography}
	book3 := books.Book{Title: "Book 3 - Author 3", Category: books.Fiction}
	book4 := books.Book{Title: "Book 3 - Author 3", Category: books.Biography}
	bookStore.AddBook(book1)
	bookStore.AddBook(book2)
	bookStore.AddBook(book3)
	bookStore.AddBook(book4)
	fmt.Println(bookStore.AmountOfBooks())
	filter := books.Filter{}
	spec := books.CategorySpecification{Category: books.Biography}
	result := filter.Filter(bookStore.Books, spec)
	for _, r := range result {
		fmt.Println(*r)
	}
}
