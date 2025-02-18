package server

import "github.com/exception-raised/src/service"

type Option func(*Server)

func WithValidationService(validationSvc *service.ValidationService) Option {
	return func(s *Server) {
		s.validationSvc = validationSvc
	}
}
