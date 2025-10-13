package teachers

import (
	"context"
	"school-management-system/internal/models"
)

type TeacherService interface {
	CREATE(context.Context, models.Teacher) (*models.Teacher, error)
}
