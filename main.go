package main

// go mod init singleresponsability

import (
	"fmt"
	"mainmodule/application"
	"mainmodule/models"
	"mainmodule/repository"
)

func main() {
	bookStoreRepository := &repository.BookStoreRepository{}
	bookStoreApp := application.BookStoreApp{BookStoreRepository: bookStoreRepository}
	bookStore := bookStoreApp.Get()

	fmt.Println("Total of books: ", bookStore.AmountOfBooks())
	biographyBooks := bookStoreApp.GetBooksByCategory(models.Biography)

	fmt.Println("Biography Books:")
	for _, r := range biographyBooks {
		fmt.Println(*r)
	}

	fmt.Println("All books:")
	for _, books := range bookStore.Books {
		fmt.Println(books)
	}
}
