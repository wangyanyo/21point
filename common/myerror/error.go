package myerror

type MyError struct {
	s string
}

func New(text string) error {
	return &MyError{text}
}

func (e *MyError) Error() string {
	return e.s
}
