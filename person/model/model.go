package model

import (
	"context"
	"github.com/joaosoft/example/person/domain"
)

func NewModel(repository domain.IRepository) domain.IModel {
	return &Model{
		repository: repository,
	}
}

func (m *Model) GetPersonById(ctx context.Context, id int) (*domain.Person, error) {
	return m.repository.GetPersonById(ctx, id)
}
