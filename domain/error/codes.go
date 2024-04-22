package error

var (
	ErrorNotFound = createError(1, "Not Found", LevelError)
)
