package infrastructure

import (
	"mainmodule/application"
	"mainmodule/repository"
)

func InitializeBookStoreApp() application.BookStoreApp {
	// Dependency injection is a software engineering technique where an object or struct receives its dependencies at compile time.
	// Hence, we need a package like "Wire" to do so
	bookStoreRepository := &repository.BookStoreRepository{}
	bookStoreApp := application.BookStoreApp{BookStoreRepository: bookStoreRepository}
	return bookStoreApp
}
