package application

import (
	"mainmodule/models"
	"mainmodule/repository"
)

type IBookStoreApp interface {
	Get() *models.BookStore
}

type BookStoreApp struct {
	BookStoreRepository repository.IBookStoreRepository
}

func (bp *BookStoreApp) Get() *models.BookStore {
	return bp.BookStoreRepository.Get()
}
