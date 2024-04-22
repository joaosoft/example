package http

import "github.com/gin-gonic/gin"

type IRegister interface {
	Register(router *gin.Engine)
}
