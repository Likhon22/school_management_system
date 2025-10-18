package repository

import (
	"context"
	"database/sql"
	"fmt"
	"school-management-system/internal/models"
	"school-management-system/pkg/utils"

	"strings"
)

type classRepo struct {
	db *sql.DB
}
type ClassRepo interface {
	Create(context.Context, models.Class) (*models.Class, error)
	Get(context.Context, map[string]string, utils.SortOption) ([]*models.Class, error)
	GetClassById(context.Context, int) (*models.Class, error)
	Update(context.Context, map[string]interface{}, map[string]bool, int) (*models.Class, error)
	Delete(ctx context.Context, id int) error
}

func NewClassRepo(db *sql.DB) ClassRepo {
	return &classRepo{
		db: db,
	}
}

func (repo *classRepo) Create(ctx context.Context, class models.Class) (*models.Class, error) {
	query := `
INSERT INTO class (name)
VALUES ($1)
RETURNING id, name, created_at, updated_at
`

	createdClass := &models.Class{}
	err := repo.db.QueryRowContext(ctx, query, class.Name).Scan(
		&createdClass.ID,
		&createdClass.Name,
		&createdClass.CreatedAt,
		&createdClass.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return createdClass, nil
}
func (repo *classRepo) Get(ctx context.Context, filters map[string]string, sort utils.SortOption) ([]*models.Class, error) {
	query := `SELECT id, name, created_at, updated_at FROM class`
	filteredQuery, args := utils.BuildFilteredQuery(query, filters, true)
	finalQuery := filteredQuery + utils.BuildSortQuery(sort)

	rows, err := repo.db.QueryContext(ctx, finalQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var class []*models.Class
	for rows.Next() {
		s := &models.Class{}
		if err := rows.Scan(
			&s.ID,
			&s.Name,
			&s.CreatedAt,
			&s.UpdatedAt,
		); err != nil {
			return nil, err
		}
		class = append(class, s)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return class, nil
}

func (repo *classRepo) GetClassById(ctx context.Context, id int) (*models.Class, error) {
	query := `SELECT id, name, created_at, updated_at FROM class WHERE id = $1`

	class := &models.Class{}
	err := repo.db.QueryRowContext(ctx, query, id).Scan(
		&class.ID,
		&class.Name,
		&class.CreatedAt,
		&class.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return class, nil
}
func (repo *classRepo) Update(ctx context.Context, data map[string]interface{}, allowedFields map[string]bool, id int) (*models.Class, error) {
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
	query := fmt.Sprintf(`UPDATE class SET %s, updated_at = NOW() WHERE id = $%d RETURNING id, name, created_at, updated_at`,
		strings.Join(setClauses, ", "), argsPos)

	class := &models.Class{}
	err := repo.db.QueryRowContext(ctx, query, args...).Scan(
		&class.ID,
		&class.Name,
		&class.CreatedAt,
		&class.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return class, nil
}
func (repo *classRepo) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM class WHERE id = $1`
	res, err := repo.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("class not found")
	}

	return nil
}
