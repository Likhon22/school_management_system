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

type studentRepo struct {
	db *sql.DB
}

type StudentRepo interface {
	Create(context.Context, models.Student) (*models.Student, error)
	Get(context.Context, map[string]string, utils.SortOption) ([]*models.Student, error)
	GetStudentById(context.Context, int) (*models.Student, error)
	Update(context.Context, map[string]interface{}, map[string]bool, int) (*models.Student, error)
	Delete(ctx context.Context, id int) error
}

func NewStudentRepo(db *sql.DB) StudentRepo {
	return &studentRepo{
		db: db,
	}

}

func (repo *studentRepo) Create(ctx context.Context, student models.Student) (*models.Student, error) {
	query := `
INSERT INTO students (
    first_name, 
    last_name, 
    email, 
    class, 
    subject
) 
VALUES ($1, $2, $3, $4, $5)
RETURNING id, first_name, last_name, email, class , created_at, updated_at
`

	createdStudent := &models.Student{}

	err := repo.db.QueryRowContext(
		ctx,
		query,
		student.FirstName,
		student.LastName,
		student.Email,
		student.Class,
	).Scan(
		&createdStudent.ID,
		&createdStudent.FirstName,
		&createdStudent.LastName,
		&createdStudent.Email,
		&createdStudent.Class,
		&createdStudent.CreatedAt,
		&createdStudent.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return createdStudent, nil
}

func (repo *studentRepo) Get(ctx context.Context, filters map[string]string, sort utils.SortOption) ([]*models.Student, error) {

	query := `SELECT id, first_name, last_name, email, class, subject, created_at, updated_at FROM students`
	filteredQuery, args := utils.BuildFilteredQuery(query, filters, true)
	finalQuery := filteredQuery + utils.BuildSortQuery(sort)
	log.Println(finalQuery)
	rows, err := repo.db.QueryContext(ctx, finalQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []*models.Student

	for rows.Next() {
		s := &models.Student{}
		if err := rows.Scan(
			&s.ID,
			&s.FirstName,
			&s.LastName,
			&s.Email,
			&s.Class,
			&s.CreatedAt,
			&s.UpdatedAt,
		); err != nil {
			return nil, err
		}
		students = append(students, s)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return students, nil
}

func (repo *studentRepo) GetStudentById(ctx context.Context, id int) (*models.Student, error) {
	query := `SELECT id, first_name, last_name, email, class, subject, created_at, updated_at FROM students WHERE id = $1`

	student := &models.Student{}
	err := repo.db.QueryRowContext(ctx, query, id).Scan(
		&student.ID,
		&student.FirstName,
		&student.LastName,
		&student.Email,
		&student.Class,
		&student.CreatedAt,
		&student.UpdatedAt,
	)

	if err == sql.ErrNoRows {

		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return student, nil
}

func (repo *studentRepo) Update(ctx context.Context, student map[string]interface{}, allowedFields map[string]bool, id int) (*models.Student, error) {
	if len(student) == 0 {
		return nil, nil

	}
	args := []interface{}{}
	argsPos := 1
	setClauses := []string{}
	for k, v := range student {

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
	query := fmt.Sprintf(`UPDATE students SET %s, updated_at = NOW() WHERE id = $%d RETURNING id, first_name, last_name, email, class, subject, created_at, updated_at`, strings.Join(setClauses, ", "), argsPos)
	var updatedStudent models.Student
	err := repo.db.QueryRowContext(ctx, query, args...).Scan(
		&updatedStudent.ID,
		&updatedStudent.FirstName,
		&updatedStudent.LastName,
		&updatedStudent.Email,
		&updatedStudent.Class,
		&updatedStudent.CreatedAt,
		&updatedStudent.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &updatedStudent, nil
}

func (repo *studentRepo) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM students WHERE id = $1`
	res, err := repo.db.ExecContext(ctx, query, id)
	if err != nil {
		return err

	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err

	}
	if rowsAffected == 0 {
		return fmt.Errorf("student not found")

	}
	return nil
}
