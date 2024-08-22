package common

import "errors"

var (
	InternalError = errors.New("internal error")
	NotFound      = errors.New("entity not found")
	BadRequest    = errors.New("invalid parameters")
)
