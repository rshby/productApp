package customError

type NotFoundError struct {
	s string
}

// function provider
func NewNotFoundError(msg string) error {
	return &NotFoundError{msg}
}

func (n *NotFoundError) Error() string {
	return n.s
}

var ErrorNotFound = &NotFoundError{}
