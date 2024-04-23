package model

import (
	"context"
	errorCodes "github.com/joaosoft/example/domain/error"
	"github.com/joaosoft/example/domain/test"
	"github.com/joaosoft/example/person/domain"
)

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

func testCaseSavePersonWithSuccess() *TestCase {
	// request
	ctx := context.Background()
	savePerson := &domain.SavePerson{
		Name: "JOAO",
		Age:  36,
	}

	// response
	successResponse := &domain.CreatedPerson{
		Id: 1,
	}

	return &TestCase{
		BaseTestCase: test.BaseTestCase{
			Test:        1,
			Description: "Saving a person with success",
			Call: test.Call{
				Arguments: []interface{}{ctx, savePerson},
				Expected:  []interface{}{successResponse, nil},
			},
		},
		Repository: test.MapCall{
			"SavePerson": test.CallList{
				test.Call{
					Arguments: []interface{}{ctx, savePerson},
					Expected:  []interface{}{successResponse, nil},
				},
			},
		},
	}
}

func testCaseSavePersonWithError() *TestCase {
	// request
	ctx := context.Background()
	savePerson := &domain.SavePerson{
		Name: "JOAO",
		Age:  36,
	}

	// response
	errorResponse := errorCodes.ErrorNotFound()

	return &TestCase{
		BaseTestCase: test.BaseTestCase{
			Test:        2,
			Description: "Saving a person with error",
			Call: test.Call{
				Arguments: []interface{}{ctx, savePerson},
				Expected:  []interface{}{nil, errorResponse},
			},
		},
		Repository: test.MapCall{
			"SavePerson": test.CallList{
				test.Call{
					Arguments: []interface{}{ctx, savePerson},
					Expected:  []interface{}{nil, errorResponse},
				},
			},
		},
	}
}
