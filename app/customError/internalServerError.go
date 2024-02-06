package customError

type InternalServerError struct {
	s string
}

// function provider
func NewInternalSeverError(msg string) error {
	return &InternalServerError{msg}
}

func (i *InternalServerError) Error() string {
	return i.s
}

var ErrorInternalServer = &InternalServerError{}
