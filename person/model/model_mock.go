package model

import (
	"context"
	"github.com/joaosoft/example/domain/person"
	"github.com/stretchr/testify/mock"
)

func NewModelMock() *ModelMock {
	return &ModelMock{}
}

type ModelMock struct {
	mock.Mock
}

func (m *ModelMock) GetPerson(ctx context.Context, id int) (*person.Person, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*person.Person), args.Error(1)
}
