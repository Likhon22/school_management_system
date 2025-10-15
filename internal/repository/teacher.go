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

type teacherRepo struct {
	db *sql.DB
}

type TeacherRepo interface {
	Create(context.Context, models.Teacher) (*models.Teacher, error)
	Get(context.Context, map[string]string, utils.SortOption) ([]*models.Teacher, error)
	GetTeacherById(context.Context, int) (*models.Teacher, error)
	GetStudentsByTeacherID(ctx context.Context, id int) (*models.TeacherWithStudents, error)
	Update(context.Context, map[string]interface{}, map[string]bool, int) (*models.Teacher, error)
	Delete(ctx context.Context, id int) error
}

func NewTeacherRepo(db *sql.DB) TeacherRepo {
	return &teacherRepo{
		db: db,
	}
}

func (repo *teacherRepo) Create(ctx context.Context, teacher models.Teacher) (*models.Teacher, error) {
	query := `
INSERT INTO teachers (
    first_name, 
    last_name, 
    email, 
    class_id, 
    subject
) 
VALUES ($1, $2, $3, $4, $5)
RETURNING id, first_name, last_name, email, class_id, subject, created_at, updated_at
`

	createdTeacher := &models.Teacher{}

	err := repo.db.QueryRowContext(
		ctx,
		query,
		teacher.FirstName,
		teacher.LastName,
		teacher.Email,
		teacher.ClassID, // fixed
		teacher.Subject,
	).Scan(
		&createdTeacher.ID,
		&createdTeacher.FirstName,
		&createdTeacher.LastName,
		&createdTeacher.Email,
		&createdTeacher.ClassID, // fixed
		&createdTeacher.Subject,
		&createdTeacher.CreatedAt,
		&createdTeacher.UpdatedAt,
	)

	if err != nil {

		return nil, err
	}

	return createdTeacher, nil
}

func (repo *teacherRepo) Get(ctx context.Context, filters map[string]string, sort utils.SortOption) ([]*models.Teacher, error) {
	query := `SELECT id, first_name, last_name, email, class_id, subject, created_at, updated_at FROM teachers`
	filteredQuery, args := utils.BuildFilteredQuery(query, filters, true)
	finalQuery := filteredQuery + utils.BuildSortQuery(sort)
	log.Println(finalQuery)
	rows, err := repo.db.QueryContext(ctx, finalQuery, args...)
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
			&t.ClassID, // fixed
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

func (repo *teacherRepo) GetTeacherById(ctx context.Context, id int) (*models.Teacher, error) {
	query := `SELECT id, first_name, last_name, email, class_id, subject, created_at, updated_at FROM teachers WHERE id = $1`

	teacher := &models.Teacher{}
	err := repo.db.QueryRowContext(ctx, query, id).Scan(
		&teacher.ID,
		&teacher.FirstName,
		&teacher.LastName,
		&teacher.Email,
		&teacher.ClassID, // fixed
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

func (repo *teacherRepo) Update(ctx context.Context, teacher map[string]interface{}, allowedFields map[string]bool, id int) (*models.Teacher, error) {
	if len(teacher) == 0 {
		return nil, nil
	}

	args := []interface{}{}
	argsPos := 1
	setClauses := []string{}
	for k, v := range teacher {
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
	query := fmt.Sprintf(`UPDATE teachers SET %s, updated_at = NOW() WHERE id = $%d RETURNING id, first_name, last_name, email, class_id, subject, created_at, updated_at`, strings.Join(setClauses, ", "), argsPos)
	var updatedTeacher models.Teacher
	err := repo.db.QueryRowContext(ctx, query, args...).Scan(
		&updatedTeacher.ID,
		&updatedTeacher.FirstName,
		&updatedTeacher.LastName,
		&updatedTeacher.Email,
		&updatedTeacher.ClassID, // fixed
		&updatedTeacher.Subject,
		&updatedTeacher.CreatedAt,
		&updatedTeacher.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &updatedTeacher, nil
}

func (repo *teacherRepo) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM teachers WHERE id = $1`
	res, err := repo.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("teacher not found")
	}
	return nil
}

func (repo *teacherRepo) GetStudentsByTeacherID(ctx context.Context, id int) (*models.TeacherWithStudents, error) {
	query := `SELECT 
    t.id AS teacher_id,
    t.first_name AS teacher_first_name,
    t.last_name AS teacher_last_name,
    t.email AS teacher_email,
    s.id AS student_id,
    s.first_name AS student_first_name,
    s.last_name AS student_last_name,
    s.email AS student_email,
    s.class_id
FROM teachers t
JOIN students s ON s.class_id = t.class_id
WHERE t.id = $1;
`

	rows, err := repo.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	teacher := &models.TeacherWithStudents{
		Students: []*models.Student{},
	}

	for rows.Next() {
		var s models.Student
		if err := rows.Scan(
			&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.Email,
			&s.ID, &s.FirstName, &s.LastName, &s.Email, &s.ClassID,
		); err != nil {
			return nil, err
		}
		teacher.Students = append(teacher.Students, &s)
	}
	teacher.TotalStudents = len(teacher.Students)

	return teacher, rows.Err()
}
