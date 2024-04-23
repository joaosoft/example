package model

import (
	"github.com/joaosoft/example/domain/test"
	"github.com/joaosoft/example/person/domain"
	"github.com/stretchr/testify/mock"
)

type Model struct {
	repository domain.IRepository
}

type ModelMock struct {
	mock.Mock
}

type TestCase struct {
	test.BaseTestCase
	Repository test.MapCall
}
