package person

type GetPersonByIDRequest struct {
	Id int `json:"id" uri:"id" validate:"required,gte=1"`
}
