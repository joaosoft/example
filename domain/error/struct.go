package error

type Level int

type errorInstance struct {
	err Error
}
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Level   Level  `json:"level"`
}
