package service

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"school-management-system/internal/config"
	"school-management-system/internal/models"
	"school-management-system/internal/repository"
	"school-management-system/pkg/utils"
	"time"
)

var (
	ErrExecNotFound    = errors.New("no exec found with that email")
	ErrPasswordInvalid = errors.New("invalid password")
	ErrSamePassword    = errors.New("same password")
)

type execService struct {
	repo repository.ExecRepo
	cnf  *config.AuthConfig
}

type ExecService interface {
	Create(ctx context.Context, exec *models.Exec) (*models.Exec, error)
	Get(ctx context.Context, filters map[string]string, sort utils.SortOption) ([]*models.Exec, error)
	GetExecById(ctx context.Context, id int) (*models.Exec, error)
	Update(ctx context.Context, fields map[string]interface{}, allowedFields map[string]bool, id int) (*models.Exec, error)
	Delete(ctx context.Context, id int) error
	Login(ctx context.Context, email, password string) (*models.ResExec, string, error)
	UpdatePassword(ctx context.Context, id int, currentPassword, newPassword string) (string, error)
	ForgetPassword(ctx context.Context, email string) (string, error)
}

func NewExecService(repo repository.ExecRepo, cnf *config.AuthConfig) ExecService {
	return &execService{
		repo: repo,
		cnf:  cnf,
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

	token, err := utils.SignedToken(exec.ID, exec.Email, exec.Username, string(exec.Role), s.cnf.JwtSecret, s.cnf.JwtExpires)
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

func (s *execService) UpdatePassword(ctx context.Context, id int, currentPassword, newPassword string) (string, error) {
	if currentPassword == newPassword {
		return "", ErrSamePassword
	}
	exec, err := s.repo.GetExecById(ctx, id)
	if err != nil {
		return "", err
	}

	if exec == nil {
		return "", ErrExecNotFound
	}

	valid, err := utils.VerifyPassword(currentPassword, exec.Password)
	if err != nil {
		return "", err
	}
	if !valid {
		return "", ErrPasswordInvalid
	}

	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return "", err
	}

	if err := s.repo.UpdatePassword(ctx, id, hashedPassword); err != nil {
		return "", err
	}
	token, err := utils.SignedToken(exec.ID, exec.Email, exec.Username, string(exec.Role), s.cnf.JwtSecret, s.cnf.JwtExpires)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *execService) ForgetPassword(ctx context.Context, email string) (string, error) {
	if email == "" {
		return "", ErrExecNotFound
	}

	user, err := s.repo.GetExecByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", ErrExecNotFound
	}

	// Calculate exact expiration timestamp
	expiryTime := time.Now().Add(s.cnf.ResetTokenExpDuration)

	// Generate secure random token
	tokenBytes := make([]byte, 32)
	_, err = rand.Read(tokenBytes)
	if err != nil {
		return "", err
	}

	// Token to send in email
	token := hex.EncodeToString(tokenBytes)

	// Hash token to store in DB
	hashToken := sha256.Sum256(tokenBytes)
	hashTokenString := hex.EncodeToString(hashToken[:])

	// Update DB with hashed token & expiry
	err = s.repo.UpdateResetToken(ctx, hashTokenString, expiryTime, user.ID)
	if err != nil {
		return "", err
	}

	// Construct reset URL
	resetURL := fmt.Sprintf("http://localhost:3000/execs/resetpassword/reset/%s", token)

	// Message for email
	message := fmt.Sprintf(
		"Forgot your password? Reset it using the following link:\n%s\nIf you didn't request this, ignore this email. This link is valid for %d minutes.",
		resetURL,
		int(s.cnf.ResetTokenExpDuration.Minutes()),
	)
	err = utils.SendMail("likhonsarker793@gmail.com", []string{user.Email}, "Your password reset link", message)
	if err != nil {
		return "", err
	}

	return "password reset link mailed to user", nil
}
