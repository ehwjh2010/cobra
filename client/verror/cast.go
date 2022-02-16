package verror

type CastError struct {
	error   error
	message string
}

func (c *CastError) Message(message string) {
	c.message = message
}

func (c CastError) Error() string {
	return c.Error() + "detail: " + c.message
}

func NewCastError(error error) *CastError {
	return &CastError{error: error}
}

//var (
//	AnyBaseCastError = NewCastError(errors.New("interface cast error"))
//)
