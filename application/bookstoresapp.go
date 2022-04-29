package application

import (
	"mainmodule/models"
	"mainmodule/repository"
)

type IBookStoreApp interface {
	Get() *models.BookStore
	GetBooksByCategory(category models.Category) *[]models.Book
}

type BookStoreApp struct {
	BookStoreRepository repository.IBookStoreRepository
}

func (bp *BookStoreApp) Get() *models.BookStore {
	return bp.BookStoreRepository.Get()
}

func (bp *BookStoreApp) GetBooksByCategory(category models.Category) []*models.Book {
	bookStore := bp.BookStoreRepository.Get()
	return bookStore.Filter(models.CategorySpecification{Category: models.Biography})
}
