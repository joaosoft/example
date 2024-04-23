package domain

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joaosoft/example/domain/http"
)

type IController interface {
	http.IRegister
	GetPersonById(ctx *gin.Context)
}

type IModel interface {
	GetPersonById(ctx context.Context, id int) (*Person, error)
}

type IRepository interface {
	GetPersonById(ctx context.Context, id int) (*Person, error)
}
