package domain

type GetPersonByIDRequest struct {
	Id int `json:"id" uri:"id" validate:"required,gte=1"`
}

type SavePersonRequest struct {
	Body struct {
		Name string `json:"name" validate:"required,gte=1"`
		Age  int    `json:"age" validate:"required,gte=1"`
	} `json:"body"`
}
