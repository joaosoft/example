package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func New(port int) *Server {
	gin.SetMode(gin.DebugMode)

	router := gin.New()
	server := &Server{
		app: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: router,
		},
		router: router,
		port:   port,
	}

	router.Use(gin.Recovery())

	return server
}

func (s *Server) Controllers(controllers ...IRegister) *Server {
	for _, controller := range controllers {
		controller.Register(s.router)
	}

	s.router.NoRoute(func(c *gin.Context) {
		errorCode := http.StatusNotFound
		c.JSON(errorCode, struct {
			Code  int    `json:"code"`
			Error string `json:"error"`
		}{
			Code:  errorCode,
			Error: http.StatusText(errorCode),
		})
	})

	return s
}

func (s *Server) Start() (err error) {
	return s.app.ListenAndServe()
}

func (s *Server) Stop() (err error) {
	return s.app.Shutdown(context.Background())
}
