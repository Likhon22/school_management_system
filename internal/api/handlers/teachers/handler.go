package teachers

import (
	"school-management-system/internal/service"
	"school-management-system/internal/validation"
)

type Handler struct {
	service   service.TeacherService
	validator *validation.Validator
}

func NewHandler(TeacherService service.TeacherService, validator *validation.Validator) *Handler {
	return &Handler{
		service:   TeacherService,
		validator: validator,
	}

}
