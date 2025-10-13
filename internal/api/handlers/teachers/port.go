package teachers

import (
	"context"
	"school-management-system/internal/models"
)

type TeacherService interface {
	CREATE(context.Context, models.Teacher) (*models.Teacher, error)
	Get(ctx context.Context, firstName, lastName string) ([]*models.Teacher, error)
	GetTeacherById(ctx context.Context, id int) (*models.Teacher, error)
}
