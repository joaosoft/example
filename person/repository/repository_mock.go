package repository

import (
	"context"
	"github.com/joaosoft/example/person/domain"
	"github.com/stretchr/testify/mock"
)

func NewRepositoryMock() *RepositoryMock {
	return &RepositoryMock{}
}

type RepositoryMock struct {
	mock.Mock
}

func (m *RepositoryMock) GetPersonById(ctx context.Context, id int) (*domain.Person, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {

		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Person), args.Error(1)
}
