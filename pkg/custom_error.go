package pkg

import "errors"

var ErrNotFound = errors.New("data not found")
var ErrBadRequest = errors.New("bad request")
var ErrInternalServerError = errors.New("internal server error")
