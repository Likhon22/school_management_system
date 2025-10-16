package service

import (
	"context"
	"school-management-system/internal/models"
	"school-management-system/internal/repository"
	"school-management-system/pkg/utils"
)

type classService struct {
	repo repository.ClassRepo
}
type ClassService interface {
	Create(context.Context, models.Class) (*models.Class, error)
	Get(context.Context, map[string]string, utils.SortOption) ([]*models.Class, error)
	GetClassById(context.Context, int) (*models.Class, error)
	Update(context.Context, map[string]interface{}, map[string]bool, int) (*models.Class, error)
	Delete(ctx context.Context, id int) error
}

func NewClassService(repo repository.ClassRepo) ClassService {
	return &classService{
		repo: repo,
	}

}

func (s *classService) Create(ctx context.Context, class models.Class) (*models.Class, error) {
	return s.repo.Create(ctx, class)
}

func (s *classService) Get(ctx context.Context, filters map[string]string, sort utils.SortOption) ([]*models.Class, error) {
	return s.repo.Get(ctx, filters, sort)
}
func (s *classService) GetClassById(ctx context.Context, id int) (*models.Class, error) {
	return s.repo.GetClassById(ctx, id)
}
func (s *classService) Update(ctx context.Context, class map[string]interface{}, allowedFields map[string]bool, id int) (*models.Class, error) {

	return s.repo.Update(ctx, class, allowedFields, id)
}
func (s *classService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
