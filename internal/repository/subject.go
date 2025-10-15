package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"school-management-system/internal/models"
	"school-management-system/pkg/utils"
	"strings"
)

type subjectRepo struct {
	db *sql.DB
}
type SubjectRepo interface {
	Create(context.Context, models.Subject) (*models.Subject, error)
	Get(context.Context, map[string]string, utils.SortOption) ([]*models.Subject, error)
	GetSubjectById(context.Context, int) (*models.Subject, error)
	Update(context.Context, map[string]interface{}, map[string]bool, int) (*models.Subject, error)
	Delete(ctx context.Context, id int) error
}

func NewSubjectRepo(db *sql.DB) SubjectRepo {
	return &subjectRepo{
		db: db,
	}
}

func (repo *subjectRepo) Create(ctx context.Context, subject models.Subject) (*models.Subject, error) {
	query := `
INSERT INTO subjects (name)
VALUES ($1)
RETURNING id, name, created_at, updated_at
`

	createdSubject := &models.Subject{}
	err := repo.db.QueryRowContext(ctx, query, subject.Name).Scan(
		&createdSubject.ID,
		&createdSubject.Name,
		&createdSubject.CreatedAt,
		&createdSubject.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return createdSubject, nil
}
func (repo *subjectRepo) Get(ctx context.Context, filters map[string]string, sort utils.SortOption) ([]*models.Subject, error) {
	query := `SELECT id, name, created_at, updated_at FROM subjects`
	filteredQuery, args := utils.BuildFilteredQuery(query, filters, true)
	finalQuery := filteredQuery + utils.BuildSortQuery(sort)
	log.Println(finalQuery)

	rows, err := repo.db.QueryContext(ctx, finalQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subjects []*models.Subject
	for rows.Next() {
		s := &models.Subject{}
		if err := rows.Scan(
			&s.ID,
			&s.Name,
			&s.CreatedAt,
			&s.UpdatedAt,
		); err != nil {
			return nil, err
		}
		subjects = append(subjects, s)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return subjects, nil
}

func (repo *subjectRepo) GetSubjectById(ctx context.Context, id int) (*models.Subject, error) {
	query := `SELECT id, name, created_at, updated_at FROM subjects WHERE id = $1`

	subject := &models.Subject{}
	err := repo.db.QueryRowContext(ctx, query, id).Scan(
		&subject.ID,
		&subject.Name,
		&subject.CreatedAt,
		&subject.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return subject, nil
}
func (repo *subjectRepo) Update(ctx context.Context, data map[string]interface{}, allowedFields map[string]bool, id int) (*models.Subject, error) {
	if len(data) == 0 {
		return nil, nil
	}

	args := []interface{}{}
	argsPos := 1
	setClauses := []string{}

	for k, v := range data {
		if allowedFields[k] {
			setClauses = append(setClauses, fmt.Sprintf("%s=$%d", k, argsPos))
			args = append(args, v)
			argsPos++
		}
	}

	if len(setClauses) == 0 {
		return nil, nil
	}

	args = append(args, id)
	query := fmt.Sprintf(`UPDATE subjects SET %s, updated_at = NOW() WHERE id = $%d RETURNING id, name, created_at, updated_at`,
		strings.Join(setClauses, ", "), argsPos)

	subject := &models.Subject{}
	err := repo.db.QueryRowContext(ctx, query, args...).Scan(
		&subject.ID,
		&subject.Name,
		&subject.CreatedAt,
		&subject.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return subject, nil
}
func (repo *subjectRepo) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM subjects WHERE id = $1`
	res, err := repo.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("subject not found")
	}

	return nil
}
