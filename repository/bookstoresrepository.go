package repository

import (
	"mainmodule/models"
)

type IBookStoreRepository interface {
	Get() *models.BookStore
}

type BookStoreRepository struct {
}

func (br *BookStoreRepository) Get() *models.BookStore {
	bookStore := models.BookStore{}
	book1 := models.Book{Product: models.Product{Name: "Book 1 - Author 1", Category: models.Adventure}, Pages: 134}
	book2 := models.Book{Product: models.Product{Name: "Book 2 - Author 2", Category: models.Biography}, Pages: 353}
	book3 := models.Book{Product: models.Product{Name: "Book 3 - Author 3", Category: models.Fiction}, Pages: 189}
	book4 := models.Book{Product: models.Product{Name: "Book 4 - Author 4", Category: models.Biography}, Pages: 230}

	bookStore.AddBook(book1)
	bookStore.AddBook(book2)
	bookStore.AddBook(book3)
	bookStore.AddBook(book4)

	return &bookStore
}
