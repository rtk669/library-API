package serverHTTP

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	handlers *Handlers
}

func NewServer(handlers *Handlers) *Server {
	return &Server{
		handlers: handlers,
	}
}

func (s *Server) StartServer() error {
	router := mux.NewRouter()

	router.Path("/books").Methods("POST").HandlerFunc(s.handlers.AddBookHandler)
	router.Path("/books").Methods("GET").HandlerFunc(s.handlers.FilteredBooksHandler)

	router.Path("/books/{title}").Methods("GET").HandlerFunc(s.handlers.GetBookHeandler)
	router.Path("/books/{title}").Methods("DELETE").HandlerFunc(s.handlers.DeleteBookHeandler)
	router.Path("/books/{title}").Methods("PATCH").HandlerFunc(s.handlers.ReadHandler)
	err := http.ListenAndServe(":9090", router)

	return err
}
