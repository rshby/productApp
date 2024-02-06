package customError

type BadRequestError struct {
	s string
}

// function provider
func NewBadRequestError(msg string) error {
	return &BadRequestError{msg}
}

func (b *BadRequestError) Error() string {
	return b.s
}

var ErrorBadRequest = &BadRequestError{}
