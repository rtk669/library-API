package serverHTTP

import (
	"API/apperror"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type DTOBook struct {
	Title  string
	Author string
	Pages  int
}

func (b DTOBook) Validate() error {
	if b.Title == "" {
		return errors.New("title is empty")
	}

	if b.Author == "" {
		return errors.New("author is empty")
	}

	if b.Pages == 0 {
		return errors.New("no pages")
	}

	return nil
}

type DTOError struct {
	Err     string
	ErrTime time.Time
}

func (e DTOError) ToString() string {
	b, err := json.MarshalIndent(e, "", "    ")
	if err != nil {
		panic(err)
	}

	return string(b)
}

func ErrorCreation(err error, w http.ResponseWriter) {
	DTOError := DTOError{
		Err:     err.Error(),
		ErrTime: time.Now(),
	}

	if errors.Is(err, apperror.NotFound) || errors.Is(err, apperror.AlreadyExist) {
		http.Error(w, DTOError.ToString(), http.StatusBadRequest)
	} else {
		http.Error(w, DTOError.ToString(), http.StatusInternalServerError)
	}

}
