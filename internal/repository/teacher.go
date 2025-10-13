package repository

import (
	"context"
	"database/sql"
	"school-management-system/internal/api/handlers/teachers"
	"school-management-system/internal/models"
	"school-management-system/pkg/utils"
)

type teacherRepo struct {
	db *sql.DB
}
type TeacherRepo interface {
	teachers.TeacherService
}

func NewTeacherRepo(db *sql.DB) TeacherRepo {
	return &teacherRepo{
		db: db,
	}

}
func (tc *teacherRepo) CREATE(ctx context.Context, teacher models.Teacher) (*models.Teacher, error) {
	query := `
INSERT INTO teachers (
    first_name, 
    last_name, 
    email, 
    class, 
    subject
) 
VALUES ($1, $2, $3, $4, $5)
RETURNING id, first_name, last_name, email, class, subject, created_at, updated_at
`

	createdTeacher := &models.Teacher{}

	err := tc.db.QueryRowContext(
		ctx,
		query,
		teacher.FirstName,
		teacher.LastName,
		teacher.Email,
		teacher.Class,
		teacher.Subject,
	).Scan(
		&createdTeacher.ID,
		&createdTeacher.FirstName,
		&createdTeacher.LastName,
		&createdTeacher.Email,
		&createdTeacher.Class,
		&createdTeacher.Subject,
		&createdTeacher.CreatedAt,
		&createdTeacher.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return createdTeacher, nil
}

func (tc *teacherRepo) Get(ctx context.Context, filters map[string]string) ([]*models.Teacher, error) {

	query := `SELECT id, first_name, last_name, email, class, subject, created_at, updated_at FROM teachers`
	filteredQuery, args := utils.BuildFilteredQuery(query, filters, true)

	rows, err := tc.db.QueryContext(ctx, filteredQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teachers []*models.Teacher

	for rows.Next() {
		t := &models.Teacher{}
		if err := rows.Scan(
			&t.ID,
			&t.FirstName,
			&t.LastName,
			&t.Email,
			&t.Class,
			&t.Subject,
			&t.CreatedAt,
			&t.UpdatedAt,
		); err != nil {
			return nil, err
		}
		teachers = append(teachers, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return teachers, nil
}

func (tc *teacherRepo) GetTeacherById(ctx context.Context, id int) (*models.Teacher, error) {
	query := `SELECT id, first_name, last_name, email, class, subject, created_at, updated_at FROM teachers WHERE id = $1`

	teacher := &models.Teacher{}
	err := tc.db.QueryRowContext(ctx, query, id).Scan(
		&teacher.ID,
		&teacher.FirstName,
		&teacher.LastName,
		&teacher.Email,
		&teacher.Class,
		&teacher.Subject,
		&teacher.CreatedAt,
		&teacher.UpdatedAt,
	)

	if err == sql.ErrNoRows {

		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return teacher, nil
}
