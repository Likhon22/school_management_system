package service

import (
	"context"
	"school-management-system/internal/models"
	"school-management-system/internal/repository"
	"school-management-system/pkg/utils"
)

type execService struct {
	repo repository.ExecRepo
}

type ExecService interface {
	Create(ctx context.Context, exec models.Exec) (*models.Exec, error)
	Get(ctx context.Context, filters map[string]string, sort utils.SortOption) ([]*models.Exec, error)
	GetExecById(ctx context.Context, id int) (*models.Exec, error)
	Update(ctx context.Context, fields map[string]interface{}, allowedFields map[string]bool, id int) (*models.Exec, error)
	Delete(ctx context.Context, id int) error
}

func NewExecService(repo repository.ExecRepo) ExecService {
	return &execService{
		repo: repo,
	}

}

func (s *execService) Create(ctx context.Context, student models.Exec) (*models.Exec, error) {
	return s.repo.Create(ctx, student)
}

func (s *execService) Get(ctx context.Context, filters map[string]string, sort utils.SortOption) ([]*models.Exec, error) {
	return s.repo.Get(ctx, filters, sort)
}
func (s *execService) GetExecById(ctx context.Context, id int) (*models.Exec, error) {
	return s.repo.GetExecById(ctx, id)
}
func (s *execService) Update(ctx context.Context, fields map[string]interface{}, allowedFields map[string]bool, id int) (*models.Exec, error) {
	return s.repo.Update(ctx, fields, allowedFields, id)
}
func (s *execService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
