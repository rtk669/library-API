package library

import (
	"API/apperror"
	"maps"
	"sync"
)

type Library struct {
	books map[string]Book
	mtx   sync.Mutex
}

type Params struct {
	Author   *string
	IsReaded *bool
}

func NewParam(a *string, b *bool) *Params {
	return &Params{
		Author:   a,
		IsReaded: b,
	}
}

func NewLibrary() *Library {
	return &Library{
		books: make(map[string]Book),
	}
}

func (l *Library) AddBook(book Book) error {
	if _, ok := l.books[book.Title]; ok {
		return apperror.AlreadyExist
	}
	l.books[book.Title] = book
	return nil
}

func (l *Library) Read(title string) (Book, error) {
	if book, ok := l.books[title]; !ok {
		return Book{}, apperror.NotFound
	} else {
		book.Read()
		l.books[title] = book
		return book, nil
	}
}

func (l *Library) GetBook(title string) (Book, error) {
	if book, ok := l.books[title]; !ok {
		return Book{}, apperror.NotFound
	} else {
		return book, nil
	}
}

func (l *Library) ListBooks() map[string]Book {
	booksCopy := make(map[string]Book)

	maps.Copy(booksCopy, l.books)

	return booksCopy
}

func (l *Library) Filtrate(params Params) map[string]Book {
	books := make(map[string]Book)

	for k, v := range l.books {
		f := true

		if params.Author != nil {
			if *params.Author != l.books[k].Author {
				f = false
			}
		}

		if params.IsReaded != nil {
			if *params.IsReaded != l.books[k].IsReaded {
				f = false
			}
		}

		if f {
			books[k] = v
		}

	}

	return books
}

func (l *Library) DeleteBook(title string) error {
	if _, ok := l.books[title]; !ok {
		return apperror.NotFound
	}

	delete(l.books, title)

	return nil
}
