package usecase

import "errors"

var todoErr = errors.New("todo")

type domainError struct {
	Code int
	Message string
}

func (e domainError) Error() string {
	return e.Message
}
