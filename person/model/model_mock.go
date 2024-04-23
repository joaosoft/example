package model

import (
	"context"
	"github.com/joaosoft/example/person/domain"
)

func NewModelMock() *ModelMock {

	return &ModelMock{}
}

func (m *ModelMock) GetPersonById(ctx context.Context, id int) (*domain.Person, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Person), args.Error(1)
}

func (m *ModelMock) SavePerson(ctx context.Context, person *domain.SavePerson) (created *domain.CreatedPerson, err error) {
	args := m.Called(ctx, person)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.CreatedPerson), args.Error(1)
}
