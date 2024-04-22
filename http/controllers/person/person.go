package person

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	httpDomain "github.com/joaosoft/example/domain/http"
	"github.com/joaosoft/example/domain/person"
	"github.com/joaosoft/example/http/controllers"
	"github.com/joaosoft/example/http/middlewares"
	"net/http"
)

type Controller struct {
	validator *validator.Validate
	model     person.IModel
}

func NewController(validator *validator.Validate, model person.IModel) person.IController {
	return &Controller{
		validator: validator,
		model:     model,
	}
}

func (c *Controller) Register(router *gin.Engine) {
	v1 := router.Group("/v1")
	v1.Use(httpDomain.PrintRequest)

	v1Persons := v1.Group("/persons")
	v1Persons.Use(middlewares.CheckExample)

	v1Persons.Handle(http.MethodGet, "/:id", c.GetPersonById)
}

func (c *Controller) GetPersonById(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	request := GetPersonByIDRequest{}
	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusInternalServerError,
			controllers.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
		return
	}

	if err := c.validator.Struct(request); err != nil {
		ctx.JSON(http.StatusBadRequest,
			controllers.ErrorResponse{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
		return
	}

	person, err := c.model.GetPersonById(ctx.Request.Context(), request.Id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			controllers.ErrorResponse{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusOK, person)
}
