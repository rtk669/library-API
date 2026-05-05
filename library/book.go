package library

import "time"

type Book struct {
	Title    string
	Author   string
	Pages    int
	IsReaded bool
	AddedAt  time.Time
	ReadedAt *time.Time
}

func NewBook(
	title string,
	author string,
	pages int,
) Book {

	return Book{
		Title:   title,
		Author:  author,
		Pages:   pages,
		AddedAt: time.Now(),
	}
}

func (b *Book) Read() {
	timeNow := time.Now()
	b.IsReaded = true
	b.ReadedAt = &timeNow
}
