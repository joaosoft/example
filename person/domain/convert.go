package domain

import controllerDomain "github.com/joaosoft/example/person/controller/domain"

func (p *Person) ToResponse() *controllerDomain.PersonResponse {
	if p == nil {
		return nil
	}

	return &controllerDomain.PersonResponse{
		Id:   p.Id,
		Name: p.Name,
		Age:  p.Age,
	}
}

func (p *CreatedPerson) ToResponse() *controllerDomain.CreatedPersonResponse {
	if p == nil {
		return nil
	}

	return &controllerDomain.CreatedPersonResponse{
		Id: p.Id,
	}
}

func (p *SavePerson) FromRequest(request *controllerDomain.SavePersonRequest) {
	if p == nil {
		return
	}

	p.Name = request.Body.Name
	p.Age = request.Body.Age
}
