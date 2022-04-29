package books

type BookStore struct {
	books []string
}

func (bs *BookStore) AddBook(book string) int {
	bs.books = append(bs.books, book)
	return len(bs.books)
}

func (bs *BookStore) RemoveBook(book string) int {
	bs.books = append(bs.books, book)
	return len(bs.books)
}
