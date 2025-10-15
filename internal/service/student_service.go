package service

import (
	"context"
	"school-management-system/internal/models"
	"school-management-system/internal/repository"
	"school-management-system/pkg/utils"
)

type studentService struct {
	repo repository.StudentRepo
}

type StudentService interface {
	Create(context.Context, models.Student) (*models.Student, error)
	Get(context.Context, map[string]string, utils.SortOption) ([]*models.Student, error)
	GetStudentById(context.Context, int) (*models.Student, error)
	Update(context.Context, map[string]interface{}, map[string]bool, int) (*models.Student, error)
	Delete(ctx context.Context, id int) error
}

func NewStudentService(repo repository.StudentRepo) StudentService {
	return &studentService{
		repo: repo,
	}

}
func (s *studentService) Create(ctx context.Context, student models.Student) (*models.Student, error) {
	return s.repo.Create(ctx, student)
}

func (s *studentService) Get(ctx context.Context, filters map[string]string, sort utils.SortOption) ([]*models.Student, error) {
	return s.repo.Get(ctx, filters, sort)
}
func (s *studentService) GetStudentById(ctx context.Context, id int) (*models.Student, error) {
	return s.repo.GetStudentById(ctx, id)
}
func (s *studentService) Update(ctx context.Context, student map[string]interface{}, allowedFields map[string]bool, id int) (*models.Student, error) {
	return s.repo.Update(ctx, student, allowedFields, id)
}
func (s *studentService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
