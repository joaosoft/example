package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/joaosoft/example/person/domain"
	"github.com/joaosoft/logger"
)

type Controller struct {
	model     domain.IModel
	validator *validator.Validate
	log       logger.ILogger
}
