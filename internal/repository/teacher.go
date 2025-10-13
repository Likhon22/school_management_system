package repository

import (
	"context"
	"database/sql"
	"school-management-system/internal/api/handlers/teachers"
	"school-management-system/internal/models"
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
