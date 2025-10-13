package teachers

import (
	"context"
	"school-management-system/internal/models"
)

type TeacherService interface {
	CREATE(context.Context, models.Teacher) (*models.Teacher, error)
	Get(context.Context, map[string]string) ([]*models.Teacher, error)
	GetTeacherById(context.Context, int) (*models.Teacher, error)
}
