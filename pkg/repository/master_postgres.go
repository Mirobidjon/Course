package repository

import (
	"fmt"
	"github.com/Mirobidjon/course"
	"github.com/jmoiron/sqlx"
	"strings"
)

type MasterPostgres struct {
	db *sqlx.DB
}

func NewMasterPostgres(db *sqlx.DB) *MasterPostgres {
	return &MasterPostgres{db: db}
}

func (r *MasterPostgres) CreateMaster(input course.MasterInput, role string) (int, error) {
	query :=
		fmt.Sprintf("INSERT INTO %s (name, username, password, role) VALUES ($1, $2, $3, $4) RETURNING id",
			masterTable)

	row := r.db.QueryRow(query, input.Name, input.Username, input.Password, role)
	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *MasterPostgres) GetMaster(username, password, role string) (int, error) {
	var id int
	query :=
		fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password=$2 AND role=$3",
			masterTable)

	err := r.db.Get(&id, query, username, password, role)

	return id, err
}

func (r *MasterPostgres) GetAllTeacherCourse() ([]course.Course, error) {
	var courseStudent []course.Course
	queryCourse := fmt.Sprintf("SELECT id, name, description, student_group  FROM %s ",
		courseTable)

	err := r.db.Select(&courseStudent, queryCourse)
	if err != nil {
		return nil, err
	}

	return courseStudent, nil
}

func (r *MasterPostgres) GetAllMaster(role string) ([]course.OutputMaster, error) {
	var masters []course.OutputMaster

	query := fmt.Sprintf("SELECT id, name, username FROM %s WHERE role = $1", masterTable)

	err := r.db.Select(&masters, query, role)
	if err != nil {
		return nil, err
	}

	return masters, err
}

func (r *MasterPostgres) GetMasterByID(role string, id int) (course.OutputMaster, error) {
	var master course.OutputMaster
	query := fmt.Sprintf("SELECT id, name, username FROM %s WHERE id = $1 AND role = $2 ", masterTable)

	err := r.db.Get(&master, query, id, role)
	return master, err
}

func (r *MasterPostgres) UpdateMaster(input course.MasterInput, id int) error {
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

	setValues := strings.Join(values, ", ")
	args = append(args, id)

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", masterTable, setValues, argID)
	_, err := r.db.Exec(query, args...)

	return err
}

// teachers have some course but director hasn't
func (r *MasterPostgres) DeleteTeacher(role string, id int) error {
	queryCourse :=
		fmt.Sprintf("DELETE FROM %s cl USING %s tl WHERE cl.id = tl.course_id AND tl.teacher_id = $1",
			courseTable, teacherCourseTable)

	_, err := r.db.Exec(queryCourse, id)
	if err != nil {
		return err
	}

	queryTeacher := fmt.Sprintf("DELETE FROM %s WHERE id = $1 AND role = $2", masterTable)
	_, err = r.db.Exec(queryTeacher, id, role)
	return err
}
