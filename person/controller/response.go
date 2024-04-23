package controller

type PersonResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type CreatedPersonResponse struct {
	Id int `json:"id"`
}
