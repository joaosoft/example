package controller

import (
	"github.com/joaosoft/example/person/domain"
)

func (p *SavePersonRequest) ToDomain() *domain.SavePerson {
	if p == nil {
		return nil
	}

	return &domain.SavePerson{
		Name: p.Body.Name,
		Age:  p.Body.Age,
	}
}
