package service

import "github.com/godcong/go-auth-manager/config"

// Service ...
type Service struct {
	config *config.Configure
	rest   *RestServer
}

var global *Service

// New ...
func New(cfg *config.Configure) *Service {
	return &Service{
		config: cfg,
		rest:   NewRestServer(cfg),
	}
}

// Start ...
func Start(cfg *config.Configure) {
	global := New(cfg)
	global.rest.Start()
}

// Stop ...
func Stop() {
	global.rest.Stop()
}
