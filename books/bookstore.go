package books

type Category int

const (
	Biography Category = iota
	Fiction
	NonFiction
	Adventure
)

type Book struct {
	Title    string
	Category Category
}

type BookStore struct {
	Books []Book
}

func (bs *BookStore) AddBook(book Book) {
	bs.Books = append(bs.Books, book)
}

func (bs *BookStore) RemoveBook(book Book) {
	bs.Books = append(bs.Books, book)
}

func (bs *BookStore) AmountOfBooks() int {
	return len(bs.Books)
}
