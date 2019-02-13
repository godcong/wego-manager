package service

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/wego-auth-manager/config"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// RestServer ...
type RestServer struct {
	*gin.Engine
	loader *RouteLoader
	config *config.Configure
	server *http.Server
	Port   string
}

// NewRestServer ...
func NewRestServer(cfg *config.Configure) *RestServer {
	s := &RestServer{
		Engine: gin.Default(),
		config: cfg,
		Port:   config.MustString(cfg.REST.Port, ":8080"),
		loader: NewRouteLoader("v0"),
	}
	return s
}

// Start ...
func (s *RestServer) Start() {
	s.loader.router(s.Engine)

	s.server = &http.Server{
		Addr:    s.Port,
		Handler: s.Engine,
	}
	go func() {
		log.Printf("Listening and serving HTTP on %s\n", s.Port)
		if err := s.server.ListenAndServe(); err != nil {
			log.Printf("Httpserver: ListenAndServe() error: %s", err)
		}
	}()

}

// Stop ...
func (s *RestServer) Stop() {
	if err := s.server.Shutdown(nil); err != nil {
		panic(err) // failure/timeout shutting down the server gracefully
	}
}
