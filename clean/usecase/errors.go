package usecase

type domainError struct {
	Code int
	Message string
}

func (e domainError) Error() string {
	return e.Message
}
