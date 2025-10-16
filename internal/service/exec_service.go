package service

import (
	"context"
	"errors"
	"school-management-system/internal/models"
	"school-management-system/internal/repository"
	"school-management-system/pkg/utils"
	"time"
)

var (
	ErrExecNotFound    = errors.New("no exec found with that email")
	ErrPasswordInvalid = errors.New("invalid password")
)

type execService struct {
	repo      repository.ExecRepo
	jwtSecret string
	jwtExpire time.Duration
}

type ExecService interface {
	Create(ctx context.Context, exec *models.Exec) (*models.Exec, error)
	Get(ctx context.Context, filters map[string]string, sort utils.SortOption) ([]*models.Exec, error)
	GetExecById(ctx context.Context, id int) (*models.Exec, error)
	Update(ctx context.Context, fields map[string]interface{}, allowedFields map[string]bool, id int) (*models.Exec, error)
	Delete(ctx context.Context, id int) error
	Login(ctx context.Context, email, password string) (*models.ResExec, string, error)
}

func NewExecService(repo repository.ExecRepo, jwtSecret string, jwtExpire time.Duration) ExecService {
	return &execService{
		repo:      repo,
		jwtSecret: jwtSecret,
		jwtExpire: jwtExpire,
	}

}

func (s *execService) Create(ctx context.Context, exec *models.Exec) (*models.Exec, error) {
	hashPassword, err := utils.HashPassword(exec.Password)
	if err != nil {

		return nil, err
	}
	exec.Password = hashPassword
	exec.Password = hashPassword
	return s.repo.Create(ctx, exec)
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

func (s *execService) Login(ctx context.Context, email, password string) (*models.ResExec, string, error) {
	exec, err := s.repo.GetExecByEmail(ctx, email)
	if err != nil {
		return nil, "", err
	}
	if exec == nil {
		return nil, "", ErrExecNotFound
	}

	valid, err := utils.VerifyPassword(password, exec.Password)
	if err != nil {
		return nil, "", err
	}
	if !valid {
		return nil, "", ErrPasswordInvalid
	}

	token, err := utils.SignedToken(exec.ID, exec.Email, exec.Username, string(exec.Role), s.jwtSecret, s.jwtExpire)
	if err != nil {
		return nil, "", err
	}

	res := &models.ResExec{
		ID:        exec.ID,
		FirstName: exec.FirstName,
		LastName:  exec.LastName,
		Email:     exec.Email,
		Username:  exec.Username,
		Role:      string(exec.Role),
	}

	return res, token, nil
}
