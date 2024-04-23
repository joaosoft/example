package domain

import "github.com/joaosoft/example/person/controller"

func (p *Person) ToResponse() *controller.PersonResponse {
	if p == nil {
		return nil
	}

	return &controller.PersonResponse{
		Id:   p.Id,
		Name: p.Name,
		Age:  p.Age,
	}
}

func (p *CreatedPerson) ToResponse() *controller.CreatedPersonResponse {
	if p == nil {
		return nil
	}

	return &controller.CreatedPersonResponse{
		Id: p.Id,
	}
}
