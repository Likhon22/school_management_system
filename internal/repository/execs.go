package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"school-management-system/internal/models"
	"school-management-system/pkg/utils"
	"strings"
	"time"
)

type execRepo struct {
	db *sql.DB
}

type ExecRepo interface {
	Create(ctx context.Context, exec *models.Exec) (*models.Exec, error)
	Get(ctx context.Context, filters map[string]string, sort utils.SortOption) ([]*models.Exec, error)
	GetExecById(ctx context.Context, id int) (*models.Exec, error)
	GetExecByEmail(ctx context.Context, email string) (*models.Exec, error)
	Update(ctx context.Context, fields map[string]interface{}, allowedFields map[string]bool, id int) (*models.Exec, error)
	Delete(ctx context.Context, id int) error
	UpdatePassword(ctx context.Context, id int, newHashedPassword string) error
	UpdateResetToken(ctx context.Context, hashTokenString string, expiry time.Time, id int) error
	ResetPassword(ctx context.Context, hashTokenString string) (*models.Exec, error)
}

func NewExecRepo(db *sql.DB) ExecRepo {
	return &execRepo{db: db}
}

// Create new exec
func (repo *execRepo) Create(ctx context.Context, exec *models.Exec) (*models.Exec, error) {
	query := `
		INSERT INTO execs (
			first_name,
			last_name,
			email,
			username,
			password,
			role
		)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, first_name, last_name, email, username, password, role, password_changed_at, password_reset_token, password_reset_token_expire, created_at, updated_at
	`

	createdExec := &models.Exec{}

	err := repo.db.QueryRowContext(
		ctx,
		query,
		exec.FirstName,
		exec.LastName,
		exec.Email,
		exec.Username,
		exec.Password,
		exec.Role,
	).Scan(
		&createdExec.ID,
		&createdExec.FirstName,
		&createdExec.LastName,
		&createdExec.Email,
		&createdExec.Username,
		&createdExec.Password,
		&createdExec.Role,
		&createdExec.PasswordChangedAt,
		&createdExec.PasswordResetToken,
		&createdExec.PasswordResetTokenExpires,
		&createdExec.CreatedAt,
		&createdExec.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return createdExec, nil
}

// Get all execs (with filters & sorting)
func (repo *execRepo) Get(ctx context.Context, filters map[string]string, sort utils.SortOption) ([]*models.Exec, error) {
	query := `SELECT id, first_name, last_name, email, username, password, role, password_changed_at, password_reset_token, password_reset_token_expire, created_at, updated_at FROM execs`
	filteredQuery, args := utils.BuildFilteredQuery(query, filters, true)
	finalQuery := filteredQuery + utils.BuildSortQuery(sort)

	rows, err := repo.db.QueryContext(ctx, finalQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var execs []*models.Exec
	for rows.Next() {
		e := &models.Exec{}
		if err := rows.Scan(
			&e.ID,
			&e.FirstName,
			&e.LastName,
			&e.Email,
			&e.Username,
			&e.Password,
			&e.Role,
			&e.PasswordChangedAt,
			&e.PasswordResetToken,
			&e.PasswordResetTokenExpires,
			&e.CreatedAt,
			&e.UpdatedAt,
		); err != nil {
			return nil, err
		}
		execs = append(execs, e)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return execs, nil
}

// Get exec by ID
func (repo *execRepo) GetExecById(ctx context.Context, id int) (*models.Exec, error) {
	query := `SELECT id, first_name, last_name, email, username, password, role, password_changed_at, password_reset_token, password_reset_token_expire, created_at, updated_at FROM execs WHERE id = $1`

	exec := &models.Exec{}
	err := repo.db.QueryRowContext(ctx, query, id).Scan(
		&exec.ID,
		&exec.FirstName,
		&exec.LastName,
		&exec.Email,
		&exec.Username,
		&exec.Password,
		&exec.Role,
		&exec.PasswordChangedAt,
		&exec.PasswordResetToken,
		&exec.PasswordResetTokenExpires,
		&exec.CreatedAt,
		&exec.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return exec, nil
}

// Update exec
func (repo *execRepo) Update(ctx context.Context, fields map[string]interface{}, allowedFields map[string]bool, id int) (*models.Exec, error) {
	if len(fields) == 0 {
		return nil, nil
	}

	args := []interface{}{}
	setClauses := []string{}
	argPos := 1

	for k, v := range fields {
		if allowedFields[k] {
			setClauses = append(setClauses, fmt.Sprintf("%s=$%d", k, argPos))
			args = append(args, v)
			argPos++
		}
	}

	if len(setClauses) == 0 {
		return nil, nil
	}

	args = append(args, id)
	query := fmt.Sprintf(`UPDATE execs SET %s, updated_at=NOW() WHERE id=$%d RETURNING id, first_name, last_name, email, username, password, role, password_changed_at, password_reset_token, password_reset_token_expire, created_at, updated_at`, strings.Join(setClauses, ", "), argPos)

	var updatedExec models.Exec
	err := repo.db.QueryRowContext(ctx, query, args...).Scan(
		&updatedExec.ID,
		&updatedExec.FirstName,
		&updatedExec.LastName,
		&updatedExec.Email,
		&updatedExec.Username,
		&updatedExec.Password,
		&updatedExec.Role,
		&updatedExec.PasswordChangedAt,
		&updatedExec.PasswordResetToken,
		&updatedExec.PasswordResetTokenExpires,
		&updatedExec.CreatedAt,
		&updatedExec.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &updatedExec, nil
}

// Delete exec
func (repo *execRepo) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM execs WHERE id = $1`
	res, err := repo.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("exec not found")
	}

	return nil
}

// login exec

func (repo *execRepo) GetExecByEmail(ctx context.Context, email string) (*models.Exec, error) {
	query := `SELECT id, first_name, last_name, email, username, password, role, password_changed_at, password_reset_token, password_reset_token_expire, created_at, updated_at FROM execs WHERE email = $1`

	exec := &models.Exec{}
	err := repo.db.QueryRowContext(ctx, query, email).Scan(
		&exec.ID,
		&exec.FirstName,
		&exec.LastName,
		&exec.Email,
		&exec.Username,
		&exec.Password,
		&exec.Role,
		&exec.PasswordChangedAt,
		&exec.PasswordResetToken,
		&exec.PasswordResetTokenExpires,
		&exec.CreatedAt,
		&exec.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return exec, nil

}

func (repo *execRepo) UpdatePassword(ctx context.Context, id int, newHashedPassword string) error {
	query := `
		UPDATE execs
		SET password = $1,
		    password_changed_at = NOW(),
		    updated_at = NOW()
		WHERE id = $2
	`

	result, err := repo.db.ExecContext(ctx, query, newHashedPassword, id)
	if err != nil {
		return fmt.Errorf("failed to update password: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return errors.New("no exec found with the given ID")
	}

	return nil
}

func (repo *execRepo) UpdateResetToken(ctx context.Context, hashTokenString string, expiry time.Time, id int) error {
	query := `
		UPDATE execs
		SET password_reset_token = $1,
		    password_reset_token_expire = $2,
		    updated_at = NOW()
		WHERE id = $3
	`

	result, err := repo.db.ExecContext(ctx, query, hashTokenString, expiry, id)
	if err != nil {
		return fmt.Errorf("failed to update reset token: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return errors.New("no exec found with the given ID")
	}

	return nil
}

func (repo *execRepo) ResetPassword(ctx context.Context, hashTokenString string) (*models.Exec, error) {
	query := `SELECT id, first_name, last_name, email, username, password, role, password_changed_at, password_reset_token, password_reset_token_expire, created_at, updated_at FROM execs WHERE password_reset_token = $1`
	exec := &models.Exec{}
	err := repo.db.QueryRowContext(ctx, query, hashTokenString).Scan(
		&exec.ID,
		&exec.FirstName,
		&exec.LastName,
		&exec.Email,
		&exec.Username,
		&exec.Password,
		&exec.Role,
		&exec.PasswordChangedAt,
		&exec.PasswordResetToken,
		&exec.PasswordResetTokenExpires,
		&exec.CreatedAt,
		&exec.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return exec, nil
}
