package main

import (
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/joaosoft/example/domain/http"
	personController "github.com/joaosoft/example/http/controllers/person"
	personModel "github.com/joaosoft/example/person/model"
	personRepository "github.com/joaosoft/example/person/repository"
)

func main() {
	var db *sql.DB
	repository := personRepository.NewRepository(db)
	model := personModel.NewModel(repository)
	validator := validator.New()
	controller := personController.NewController(validator, model)

	if err := http.
		New(8081).
		Controllers(controller).
		Start(); err != nil {
		panic(err)
	}
}
