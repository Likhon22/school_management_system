package service

import (
	"context"
	"school-management-system/internal/models"
	"school-management-system/internal/repository"
	"school-management-system/pkg/utils"
)

type TeacherService interface {
	Create(context.Context, models.Teacher) (*models.Teacher, error)
	Get(context.Context, map[string]string, utils.SortOption) ([]*models.Teacher, error)
	GetTeacherById(context.Context, int) (*models.Teacher, error)
	Update(context.Context, map[string]interface{}, map[string]bool, int) (*models.Teacher, error)
	Delete(ctx context.Context, id int) error
}

type teacherService struct {
	repo repository.TeacherRepo
}

func NewTeacherService(repo repository.TeacherRepo) TeacherService {
	return &teacherService{
		repo: repo,
	}

}

func (s *teacherService) Create(ctx context.Context, teacher models.Teacher) (*models.Teacher, error) {
	return s.repo.Create(ctx, teacher)
}

func (s *teacherService) Get(ctx context.Context, filters map[string]string, sort utils.SortOption) ([]*models.Teacher, error) {
	return s.repo.Get(ctx, filters, sort)
}
func (s *teacherService) GetTeacherById(ctx context.Context, id int) (*models.Teacher, error) {
	return s.repo.GetTeacherById(ctx, id)
}
func (s *teacherService) Update(ctx context.Context, teacher map[string]interface{}, allowedFields map[string]bool, id int) (*models.Teacher, error) {
	return s.repo.Update(ctx, teacher, allowedFields, id)
}
func (s *teacherService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
