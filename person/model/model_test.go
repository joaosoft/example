package model

import (
	"context"
	"github.com/joaosoft/example/person/domain"
	"github.com/joaosoft/example/person/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGetPersonById
func TestGetPersonById(t *testing.T) {
	testCases := []*TestCase{
		testCaseGetPersonByIdWithSuccess(),
		testCaseGetPersonByIdWithError(),
	}

	for _, test := range testCases {
		test.Log(t)

		// repository
		repo := repository.NewRepositoryMock()
		test.Repository.Setup(repo)

		// model
		// TODO: review the logger mock
		model := NewModel(repo, nil)

		// model
		result, err := model.GetPersonById(test.Arguments[0].(context.Context), test.Arguments[1].(int))

		assert.Equal(t, test.Expected[1] == nil, err == nil)    // check nil error
		assert.Equal(t, test.Expected[0] == nil, result == nil) // check nil result
		if test.Expected[0] != nil {
			assert.Equal(t, test.Expected[0], result) // check result object
		}
	}
}

// TestSavePerson
func TestSavePerson(t *testing.T) {
	testCases := []*TestCase{
		testCaseSavePersonWithSuccess(),
		testCaseSavePersonWithError(),
	}

	for _, test := range testCases {
		test.Log(t)

		// repository
		repo := repository.NewRepositoryMock()
		test.Repository.Setup(repo)

		// model
		// TODO: review the logger mock
		model := NewModel(repo, nil)

		// model
		result, err := model.SavePerson(test.Arguments[0].(context.Context), test.Arguments[1].(*domain.SavePerson))

		assert.Equal(t, test.Expected[1] == nil, err == nil)    // check nil error
		assert.Equal(t, test.Expected[0] == nil, result == nil) // check nil result
		if test.Expected[0] != nil {
			assert.Equal(t, test.Expected[0], result) // check result object
		}
	}
}
