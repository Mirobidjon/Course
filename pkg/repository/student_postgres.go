package repository

import (
	"fmt"
	"github.com/Mirobidjon/course"
	"github.com/jmoiron/sqlx"
	"strings"
)

type AuthStudentPostgres struct {
	db *sqlx.DB
}

func NewAuthStudentPostgres(db *sqlx.DB) *AuthStudentPostgres {
	return &AuthStudentPostgres{db: db}
}

func (r AuthStudentPostgres) CreateStudent(input course.SignUpStudent) (int, error) {

	query :=
		fmt.Sprintf("INSERT INTO %s (name, username, password, groups, role) VALUES ($1, $2, $3, $4, $5) RETURNING id",
			studentTable)

	row := r.db.QueryRow(query, input.Name, input.Username, input.Password, input.Group, "student")
	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r AuthStudentPostgres) GetStudent(username, password string) (int, error) {
	var id int
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password=$2", studentTable)

	err := r.db.Get(&id, query, username, password)

	return id, err
}

func (r *AuthStudentPostgres) GetAllCourse(id int) ([]course.Course, error) {
	var group int
	query := fmt.Sprintf("SELECT groups FROM %s WHERE id = $1 ", studentTable)
	err := r.db.Get(&group, query, id)
	if err != nil {
		return nil, err
	}

	var courseStudent []course.Course
	queryCourse := fmt.Sprintf("SELECT id, name, description, student_group  FROM %s WHERE student_group = $1",
		courseTable)

	err = r.db.Select(&courseStudent, queryCourse, group)
	if err != nil {
		return nil, err
	}

	return courseStudent, nil
}

func (r *AuthStudentPostgres) GetAllStudents() ([]course.GetStudents, error) {
	var allStudent []course.GetStudents
	query := fmt.Sprintf("SELECT id, name, username, groups FROM %s", studentTable)

	err := r.db.Select(&allStudent, query)
	if err != nil {
		return nil, err
	}

	return allStudent, nil
}

func (r *AuthStudentPostgres) GetStudentByID(id int) (course.GetStudents, error) {
	var studentByID course.GetStudents
	query := fmt.Sprintf("SELECT id, name, username, groups FROM %s WHERE id = $1", studentTable)

	err := r.db.Get(&studentByID, query, id)
	if err != nil {
		return studentByID, err
	}

	return studentByID, nil
}

func (r *AuthStudentPostgres) UpdateStudent(input course.UpdateStudent, id int) error {
	values := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if input.Name != "" {
		values = append(values, fmt.Sprintf("name=$%d", argID))
		args = append(args, input.Name)
		argID++
	}

	if input.Username != "" {
		values = append(values, fmt.Sprintf("username=$%d", argID))
		args = append(args, input.Username)
		argID++
	}

	if input.Password != "" {
		values = append(values, fmt.Sprintf("password=$%d", argID))
		args = append(args, input.Password)
		argID++
	}

	setQuery := strings.Join(values, ", ")
	args = append(args, id)
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d",
		studentTable, setQuery, argID)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *AuthStudentPostgres) DeleteStudent(id int) error {
	queryStudentBook := fmt.Sprintf("DELETE FROM %s bl USING %s sl WHERE sl.book_id = bl.id AND sl.student_id = $1", bookTable, studentBookTable)

	_, err := r.db.Exec(queryStudentBook, id)
	if err != nil {
		return err
	}

	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", studentTable)
	_, err = r.db.Exec(query, id)
	return err
}

func (r *AuthStudentPostgres) GetCourseByID(courseID, studentID int) (course.Course, error) {
	var group int
	queryGroup := fmt.Sprintf("SELECT groups FROM %s WHERE id = $1", studentTable)
	err := r.db.Get(&group, queryGroup, studentID)
	if err != nil {
		return course.Course{}, err
	}

	var courseByID course.Course
	query :=
		fmt.Sprintf("SELECT id, name, description, student_group FROM %s WHERE id = $1 AND student_group = $2 ",
			courseTable)

	err = r.db.Get(&courseByID, query, courseID, group)
	return courseByID, err
}
