package error

type Level int

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Level   Level  `json:"level"`
}
