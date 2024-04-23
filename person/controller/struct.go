package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/joaosoft/example/person/domain"
)

type Controller struct {
	validator *validator.Validate
	model     domain.IModel
}
