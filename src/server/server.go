package server

import (
	"context"
	"net/http"
	"time"

	"github.com/exception-raised/src/service"
	"github.com/gorilla/mux"
)

type Server struct {
	router           *mux.Router
	http             *http.Server
	httpreadtimeout  time.Duration
	httpwritetimeout time.Duration

	validationSvc *service.ValidationService
}

func NewServer(opts ...Option) *Server {
	s := &Server{
		router:           mux.NewRouter(),
		httpreadtimeout:  5 * time.Second,
		httpwritetimeout: 10 * time.Second,
	}

	for _, opt := range opts {
		opt(s)
	}

	if s.validationSvc == nil {
		panic("validation service is required")
	}

	return s
}

func (s *Server) ListenAndServe(listenAddr string) error {
	s.http = &http.Server{
		Addr:         listenAddr,
		Handler:      s.router,
		ReadTimeout:  s.httpreadtimeout,
		WriteTimeout: s.httpwritetimeout,
	}

	return s.http.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.http.Shutdown(ctx)
}
