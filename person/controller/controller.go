package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	errorCodes "github.com/joaosoft/example/domain/error"
	httpDomain "github.com/joaosoft/example/domain/http"
	"github.com/joaosoft/example/http/middlewares"
	"github.com/joaosoft/example/person/domain"
	"github.com/joaosoft/logger"
	"net/http"
)

func NewController(model domain.IModel, validator *validator.Validate, log logger.ILogger) domain.IController {
	return &Controller{
		model:     model,
		validator: validator,
	}
}

func (c *Controller) Register(router *gin.Engine) {
	v1 := router.Group("/v1")
	v1.Use(httpDomain.PrintRequest)

	v1Persons := v1.Group("/persons")
	v1Persons.Use(middlewares.CheckExample)

	v1Persons.Handle(http.MethodGet, "/:id", c.GetPersonById)
	v1Persons.Handle(http.MethodPost, "", c.SavePerson)
}

func (c *Controller) GetPersonById(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	request := GetPersonByIDRequest{}
	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest,
			errorCodes.Error{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
				Level:   errorCodes.LevelError,
			})
		return
	}

	if err := c.validator.Struct(request); err != nil {
		ctx.JSON(http.StatusBadRequest,
			errorCodes.Error{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
				Level:   errorCodes.LevelError,
			})
		return
	}

	person, err := c.model.GetPersonById(ctx.Request.Context(), request.Id)
	if err != nil {
		if errors.Is(err, &errorCodes.Error{}) {
			ctx.JSON(http.StatusInternalServerError, err)
		} else {
			err := &errorCodes.Error{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
				Level:   errorCodes.LevelError,
			}
			ctx.JSON(http.StatusInternalServerError, err)
		}
		return
	}

	ctx.JSON(http.StatusOK, person.ToResponse())
}

func (c *Controller) SavePerson(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	request := SavePersonRequest{}
	if err := ctx.ShouldBindJSON(&request.Body); err != nil {
		ctx.JSON(http.StatusBadRequest,
			errorCodes.Error{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
				Level:   errorCodes.LevelError,
			})
		return
	}

	if err := c.validator.Struct(request); err != nil {
		ctx.JSON(http.StatusBadRequest,
			errorCodes.Error{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
				Level:   errorCodes.LevelError,
			})
		return
	}

	created, err := c.model.SavePerson(ctx.Request.Context(), request.ToDomain())
	if err != nil {
		if errors.Is(err, &errorCodes.Error{}) {
			ctx.JSON(http.StatusInternalServerError, err)
		} else {
			err := &errorCodes.Error{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
				Level:   errorCodes.LevelError,
			}
			ctx.JSON(http.StatusInternalServerError, err)
		}
		return
	}

	ctx.JSON(http.StatusOK, created.ToResponse())
}
