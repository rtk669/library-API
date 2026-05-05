package apperror

import "errors"

var NotFound error = errors.New("book not found")
var AlreadyExist error = errors.New("book already exist")
