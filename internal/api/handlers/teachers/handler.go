package teachers

import "school-management-system/internal/validation"

type Handler struct {
	service   TeacherService
	validator *validation.Validator
}

func NewHandler(TeacherService TeacherService, validator *validation.Validator) *Handler {
	return &Handler{
		service:   TeacherService,
		validator: validator,
	}

}
