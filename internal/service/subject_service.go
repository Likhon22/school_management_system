package service

import (
	"context"
	"school-management-system/internal/models"
	"school-management-system/internal/repository"
	"school-management-system/pkg/utils"
)

type subjectService struct {
	repo repository.SubjectRepo
}
type SubjectService interface {
	Create(context.Context, models.Subject) (*models.Subject, error)
	Get(context.Context, map[string]string, utils.SortOption) ([]*models.Subject, error)
	GetSubjectById(context.Context, int) (*models.Subject, error)
	Update(context.Context, map[string]interface{}, map[string]bool, int) (*models.Subject, error)
	Delete(ctx context.Context, id int) error
}

func NewSubjectService(repo repository.SubjectRepo) SubjectService {
	return &subjectService{
		repo: repo,
	}

}

func (s *subjectService) Create(ctx context.Context, subject models.Subject) (*models.Subject, error) {
	return s.repo.Create(ctx, subject)
}

func (s *subjectService) Get(ctx context.Context, filters map[string]string, sort utils.SortOption) ([]*models.Subject, error) {
	return s.repo.Get(ctx, filters, sort)
}
func (s *subjectService) GetSubjectById(ctx context.Context, id int) (*models.Subject, error) {
	return s.repo.GetSubjectById(ctx, id)
}
func (s *subjectService) Update(ctx context.Context, subject map[string]interface{}, allowedFields map[string]bool, id int) (*models.Subject, error) {
	return s.repo.Update(ctx, subject, allowedFields, id)
}
func (s *subjectService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
