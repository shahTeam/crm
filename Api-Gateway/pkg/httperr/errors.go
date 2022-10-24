package httperr

import "errors"

var (
	ErrBadRequest = errors.New("bad request")
	ErrInternal = errors.New("Internal server error")
	ErrNotFound = errors.New("data is not found")
	ErrUnauthorized = errors.New("unauthorized")
	ErrForbidden = errors.New("request forbidden")
)

