package example

import (
	"github.com/go-playground/validator/v10"
	"github.com/joaosoft/dbr"
	example "github.com/joaosoft/example/domain/config"
	"github.com/joaosoft/example/domain/http"
	personController "github.com/joaosoft/example/person/controller"
	personModel "github.com/joaosoft/example/person/model"
	personRepository "github.com/joaosoft/example/person/repository"
	"github.com/joaosoft/logger"
)

type Example struct {
	logger    logger.ILogger
	db        *dbr.Dbr
	validator *validator.Validate
	config    *example.ExampleConfig
}

// NewExample ...
func NewExample(options ...ExampleOption) (service *Example, err error) {
	service = &Example{}
	service.Reconfigure(options...)

	if service.config == nil {
		if service.config, err = example.NewConfig(); err != nil {
			return nil, err
		}
	}

	if service.logger == nil {
		logLevel, _ := logger.ParseLevel(service.config.Log.Level)
		service.logger = logger.NewLogDefault("example", logLevel)
	}

	if service.db == nil {
		if service.db, err = dbr.New(dbr.WithConfiguration(service.config.DbrConfig)); err != nil {
			return nil, err
		}
	}

	if service.validator == nil {
		service.validator = validator.New()
	}

	repository := personRepository.NewRepository(service.db, service.logger)
	model := personModel.NewModel(repository, service.logger)
	controller := personController.NewController(model, service.validator, service.logger)

	if err = http.New(service.config.Port).
		Controllers(controller).
		Start(); err != nil {
		return nil, err
	}

	return service, nil
}
