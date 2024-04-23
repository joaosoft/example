package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	app    *http.Server
	router *gin.Engine
	port   int
}
