package repository

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Mirobidjon/course"
	"github.com/jmoiron/sqlx"
)

type CoursePostgres struct {
	db *sqlx.DB
}

func NewCoursePostgres(db *sqlx.DB) *CoursePostgres {
	return &CoursePostgres{db: db}
}

func (r *CoursePostgres) CreateCourse(input course.InputCourse, teacherID int) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	js, err := json.Marshal(input.FileUrl)
	if err != nil {
		return 0, err
	}

	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, description, student_group, file_url) VALUES ($1, $2, $3, $4) RETURNING id",
		courseTable)

	row := tx.QueryRow(query, input.Name, input.Description, input.Group, js)
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}

	tcQuery := fmt.Sprintf("INSERT INTO %s (teacher_id, course_id) VALUES ($1, $2) ", teacherCourseTable)
	_, err = tx.Exec(tcQuery, teacherID, id)
	if err != nil {
		return 0, err
	}

	return id, tx.Commit()
}

func (r *CoursePostgres) DeleteCourse(courseID, teacherID int) error {
	query := fmt.Sprintf("DELETE FROM %s cl USING %s tl WHERE tl.course_id = cl.id AND tl.teacher_id = $1 AND tl.course_id = $2",
		courseTable, teacherCourseTable)

	_, err := r.db.Exec(query, teacherID, courseID)
	return err
}

func (r *CoursePostgres) UpdateCourse(name, description string, courseID, teacherID int) error {
	values := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1
	if name != "" {
		values = append(values, fmt.Sprintf("name=$%d", argID))
		args = append(args, name)
		argID++
	}

	if description != "" {
		values = append(values, fmt.Sprintf("description=$%d", argID))
		args = append(args, description)
		argID++
	}

	setValues := strings.Join(values, ", ")

	query :=
		fmt.Sprintf(
			"UPDATE %s cl SET %s FROM %s tl WHERE tl.course_id = cl.id AND tl.teacher_id = $%d AND tl.course_id = $%d",
			courseTable, setValues, teacherCourseTable, argID, argID+1)
	args = append(args, teacherID, courseID)
	_, err := r.db.Exec(query, args...)

	return err
}

func (r *CoursePostgres) GetAllCourse(teacherID int) ([]course.Course, error) {
	var teacherCourse []course.Course
	query := fmt.Sprintf("SELECT cl.id, cl.name, cl.description, cl.student_group FROM %s cl  INNER JOIN %s tl on tl.course_id = cl.id  WHERE  tl.teacher_id = $1",
		courseTable, teacherCourseTable)

	err := r.db.Select(&teacherCourse, query, teacherID)
	if err != nil {
		return nil, err
	}

	return teacherCourse, err
}

func (r *CoursePostgres) GetTeacherCourse(teacherID, courseID int) (course.Course, error) {
	var courseByID course.Course
	query := fmt.Sprintf(`SELECT cl.id, cl.name, cl.description, cl.student_group 
		FROM %s cl INNER JOIN %s tl on cl.id = tl.course_id WHERE tl.teacher_id = $1 AND tl.course_id = $2`,
		courseTable, teacherCourseTable)

	err := r.db.Get(&courseByID, query, teacherID, courseID)

	return courseByID, err
}

func (r *CoursePostgres) GetCourse(courseID int) (course.Course, error) {
	var courseByID course.Course
	query := fmt.Sprintf("SELECT id, name, description, student_group FROM %s WHERE id = $1",
		courseTable)

	err := r.db.Get(&courseByID, query, courseID)

	return courseByID, err
}
