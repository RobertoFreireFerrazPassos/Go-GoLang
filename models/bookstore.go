package models

type Category int

const (
	Biography Category = iota
	Fiction
	NonFiction
	Adventure
)

type Book struct {
	Product Product
	Pages   int
}

type BookStore struct {
	Books []Book
}

func (bs *BookStore) AddBook(book Book) {
	bs.Books = append(bs.Books, book)
}

func (bs *BookStore) AmountOfBooks() int {
	return len(bs.Books)
}

func (bs *BookStore) Filter(spec Specification) []*Book {
	result := make([]*Book, 0)
	for i, b := range bs.Books {
		if spec.IsSatisfied(&b.Product) {
			result = append(result, &bs.Books[i])
		}
	}
	return result
}
