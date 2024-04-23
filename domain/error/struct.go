package error

type Level int

const (
	// Level Fatal
	LevelFatal Level = iota
	// Level Error
	LevelError
	// Level Warning
	LevelWarning
	// Level Info
	LevelInfo
	// Level Debug
	LevelDebug
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Level   Level  `json:"level"`
}

func (e Error) Error() string {
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
