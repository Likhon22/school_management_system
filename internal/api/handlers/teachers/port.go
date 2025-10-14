package teachers

import (
	"context"
	"school-management-system/internal/models"
	"school-management-system/pkg/utils"
)

type TeacherService interface {
	CREATE(context.Context, models.Teacher) (*models.Teacher, error)
	Get(context.Context, map[string]string, utils.SortOption) ([]*models.Teacher, error)
	GetTeacherById(context.Context, int) (*models.Teacher, error)
	Update(context.Context, map[string]interface{}, map[string]bool, int) (*models.Teacher, error)
	Delete(ctx context.Context, id int) error
}
