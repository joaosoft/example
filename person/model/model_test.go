package model

import (
	"context"
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
		model := NewModelMock()

		// model
		result, err := model.GetPerson(test.Arguments[0].(context.Context), test.Arguments[1].(int))

		assert.Equal(t, test.Expected[1] == nil, err == nil)    // check nil error
		assert.Equal(t, test.Expected[0] == nil, result == nil) // check nil result
		if test.Expected[0] != nil {
			assert.Equal(t, test.Expected[0], result) // check result object
		}
	}
}
