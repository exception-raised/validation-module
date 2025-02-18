package service

import "github.com/go-playground/validator/v10"

type ValidationService struct {
	validator *validator.Validate
}

func NewValidationService() *ValidationService {
	return &ValidationService{
		validator: validator.New(),
	}
}
