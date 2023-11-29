package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/nikitanovikovdev/astrolog/api/handler"
	"github.com/nikitanovikovdev/astrolog/processor"
	"log"
	"net/http"
	"os"
)

type Server struct {
	httpServer    http.Server
	engine        *gin.Engine
	apodProcessor *processor.APODProcessor
}

func New(APODProcessor *processor.APODProcessor) *Server {
	engine := gin.New()

	return &Server{
		httpServer: http.Server{
			Addr:    os.Getenv("SERVER_IP") + ":" + os.Getenv("SERVER_PORT"),
			Handler: engine,
		},
		engine:        engine,
		apodProcessor: APODProcessor,
	}
}

func (s *Server) InitRoutes() {
	s.engine.Group("/").GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	APODHandler := handler.NewAPODHandler(s.apodProcessor)

	content := s.engine.Group("content")
	content.GET("/by-date", APODHandler.ByDate)
	content.GET("/all", APODHandler.AllContent)
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	log.Println("Stopping a server")
	return s.httpServer.Shutdown(ctx)
}
