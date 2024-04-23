package error

func (e *Error) Error() string {
	return e.Message
}

func createError(code int, message string, level Level) func() *Error {
	return func() *Error {
		return &Error{
			Code:    code,
			Message: message,
			Level:   level,
		}
	}
}
