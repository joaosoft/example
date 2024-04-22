package model

import (
	"context"
	"github.com/joaosoft/example/domain/person"
)

type Model struct {
	repository person.IRepository
}

func NewModel(repository person.IRepository) person.IModel {
	return &Model{
		repository: repository,
	}
}

func (m *Model) GetPersonById(ctx context.Context, id int) (*person.Person, error) {
	return m.repository.GetPersonById(ctx, id)
}
