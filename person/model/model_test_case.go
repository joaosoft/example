package model

import (
	"context"
	errorCodes "github.com/joaosoft/example/domain/error"
	"github.com/joaosoft/example/domain/test"
	"github.com/joaosoft/example/person/domain"
)

type TestCase struct {
	test.BaseTestCase
	Repository test.MapCall
}

func testCaseGetPersonByIdWithSuccess() *TestCase {
	// request
	ctx := context.Background()
	idPerson := 1

	// response
	successResponse := &domain.Person{
		Id:   1,
		Name: "JOAO",
		Age:  36,
	}

	return &TestCase{
		BaseTestCase: test.BaseTestCase{
			Test:        1,
			Description: "Getting person by id with success",
			Call: test.Call{
				Arguments: []interface{}{ctx, idPerson},
				Expected:  []interface{}{successResponse, nil},
			},
		},
		Repository: test.MapCall{
			"GetPersonById": test.CallList{
				test.Call{
					Arguments: []interface{}{ctx, idPerson},
					Expected:  []interface{}{successResponse, nil},
				},
			},
		},
	}
}

func testCaseGetPersonByIdWithError() *TestCase {
	// request
	ctx := context.Background()
	idPerson := 1

	errorResponse := errorCodes.ErrorNotFound()

	return &TestCase{
		BaseTestCase: test.BaseTestCase{
			Test:        2,
			Description: "Getting person by id with error",
			Call: test.Call{
				Arguments: []interface{}{ctx, idPerson},
				Expected:  []interface{}{nil, errorResponse},
			},
		},
		Repository: test.MapCall{
			"GetPersonById": test.CallList{
				test.Call{
					Arguments: []interface{}{ctx, idPerson},
					Expected:  []interface{}{nil, errorResponse},
				},
			},
		},
	}
}
