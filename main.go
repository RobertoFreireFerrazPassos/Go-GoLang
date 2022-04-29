package main

// go mod init singleresponsability

import (
	"fmt"
	"mainmodule/books"
)

func main() {
	bookStore := books.BookStore{}
	book1 := books.Book{Product: books.Product{Name: "Book 1 - Author 1", Category: books.Adventure}, Pages: 134}
	book2 := books.Book{Product: books.Product{Name: "Book 2 - Author 2", Category: books.Biography}, Pages: 353}
	book3 := books.Book{Product: books.Product{Name: "Book 3 - Author 3", Category: books.Fiction}, Pages: 189}
	book4 := books.Book{Product: books.Product{Name: "Book 3 - Author 3", Category: books.Biography}, Pages: 230}

	bookStore.AddBook(book1)
	bookStore.AddBook(book2)
	bookStore.AddBook(book3)
	bookStore.AddBook(book4)
	fmt.Println("Total of books: ", bookStore.AmountOfBooks())

	result := bookStore.Filter(books.CategorySpecification{Category: books.Biography})

	fmt.Println("Biography Books:")
	for _, r := range result {
		fmt.Println(*r)
	}

	fmt.Println("All books:")
	for _, books := range bookStore.Books {
		fmt.Println(books)
	}
}
