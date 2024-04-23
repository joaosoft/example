package model

import (
	"context"
	"github.com/joaosoft/example/person/domain"
	"github.com/joaosoft/logger"
)

func NewModel(repository domain.IRepository, log logger.ILogger) domain.IModel {
	return &Model{
		repository: repository,
		log:        log,
	}
}

func (m *Model) GetPersonById(ctx context.Context, id int) (*domain.Person, error) {
	return m.repository.GetPersonById(ctx, id)
}

func (m *Model) SavePerson(ctx context.Context, person *domain.SavePerson) (id int, err error) {
	return m.repository.SavePerson(ctx, person)
}
