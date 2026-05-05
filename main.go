package main

import (
	library "API/Library"
	"API/serverHTTP"
)

func main() {
	library := library.NewLibrary()
	handlers := serverHTTP.NewHTTPHandlers(library)
	s := serverHTTP.NewServer(handlers)

	s.StartServer()
}
