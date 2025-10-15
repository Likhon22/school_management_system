package class

import (
	"school-management-system/internal/service"
	"school-management-system/internal/validation"
)

type Handler struct {
	service   service.ClassService
	validator *validation.Validator
}

func NewHandler(service service.ClassService, validator *validation.Validator) *Handler {
	return &Handler{
		service:   service,
		validator: validator,
	}
}
