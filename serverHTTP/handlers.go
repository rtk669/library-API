package serverHTTP

import (
	library "API/Library"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Handlers struct {
	library *library.Library
}

func NewHTTPHandlers(library *library.Library) *Handlers {
	return &Handlers{
		library: library,
	}
}

/*
CONTRACT:

	pattern: /books
	method: POST
	info: json in body

	succed:
		head: 201
		body: json book

	failed:
		head: 400 409 500
		body: json error with time
*/
func (h *Handlers) AddBookHandler(w http.ResponseWriter, r *http.Request) {
	var DTOBook DTOBook
	err := json.NewDecoder(r.Body).Decode(&DTOBook)

	if err != nil {
		ErrorCreation(err, w)
		return
	}
	if err := DTOBook.Validate(); err != nil {
		ErrorCreation(err, w)
		return
	}
	book := library.NewBook(DTOBook.Title, DTOBook.Author, DTOBook.Pages)
	if err := h.library.AddBook(book); err != nil {
		ErrorCreation(err, w)
		return
	}
	b, err := json.MarshalIndent(book, "", "    ")

	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)

	w.Write(b)
}

/*
CONTRACT:

	pattern: /books/{title}
	method: GET
	info: in pattern

	succed:
		head: 200
		body: json book
	failed:
		head: 400 404 500
		body: json error with time
*/
func (h *Handlers) GetBookHeandler(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]

	book, err := h.library.GetBook(title)

	if err != nil {
		ErrorCreation(err, w)
		return
	}

	b, err := json.MarshalIndent(book, "", "    ")

	if err != nil {
		ErrorCreation(err, w)
		return
	}
	w.WriteHeader(http.StatusOK)

	w.Write(b)
}

/*
CONTRACT:

	pattern: /books
	method: GET
	info: -

	succed:
		head: 200
		body: json with map of books
	failed:
		head: 400
		body: json error with time
*/
func (h *Handlers) ListBooksHAndler(w http.ResponseWriter, r *http.Request) {
	books := h.library.ListBooks()

	b, err := json.MarshalIndent(books, "", "    ")

	if err != nil {
		ErrorCreation(err, w)
		return
	}
	w.Write(b)
}

/*
CONTRACT:

	pattern: /books/{title}
	method: PATCH
	info: in pattern

	succed:
		head: 200
		body: json book
	failed:
		head: 400 404 500
		body: json error with time
*/
func (h *Handlers) ReadHandler(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]

	book, err := h.library.Read(title)

	if err != nil {
		ErrorCreation(err, w)
		return
	}

	b, err := json.MarshalIndent(book, "", "    ")

	w.Write(b)
}

/*
CONTRACT:

	pattern: /books?{filter}={value}
	method: GET
	info: in pattern

	succed:
		head: 200
		body: json with filtered books
	failed:
		head: 400 500
		body: json errors with time
*/
func (h *Handlers) FilteredBooksHandler(w http.ResponseWriter, r *http.Request) {
	var params library.Params

	query := r.URL.Query()

	author, isReaded := query.Get("author"), query.Get("isReaded")

	if author != "" {
		params.Author = &author
	}

	if isReaded != "" {
		parsedIsReaded, err := strconv.ParseBool(isReaded)
		if err != nil {
			ErrorCreation(err, w)
			return
		}
		params.IsReaded = &parsedIsReaded
	}

	b, err := json.MarshalIndent(h.library.Filtrate(params), "", "    ")

	if err != nil {
		ErrorCreation(err, w)
		return
	}

	w.Write(b)

}

/*
CONTRACT:

	pattern: /books/{title}
	method: DELETE
	info: in pattern

	succed:
		head: 204
		body: -
	failed:
		head: ...
		body: json error with time
*/
func (h *Handlers) DeleteBookHeandler(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]

	if err := h.library.DeleteBook(title); err != nil {
		ErrorCreation(err, w)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
