package repository

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Mirobidjon/course"
	"github.com/jmoiron/sqlx"
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
	queryCourse := fmt.Sprintf("SELECT id, name, description, student_group, file_url  FROM %s WHERE student_group = $1 ",
		courseTable)

	rows, err := r.db.Queryx(queryCourse, group)
	if err != nil {
		return nil, err

	}

	for rows.Next() {
		var (
			course   course.Course
			file_url string
		)

		err := rows.Scan(&course.ID, &course.Name, &course.Description, &course.Student_group, &file_url)
		if err != nil {
			return nil, err
		}

		course.File_url = make(map[int]string)
		err = json.Unmarshal([]byte(file_url), &course.File_url)
		if err != nil {
			return nil, err
		}

		var mp = make(map[int]string)
		for i := range course.File_url {
			if i == id {
				mp[i] = course.File_url[i]
			}
		}

		course.File_url = mp
		courseStudent = append(courseStudent, course)
	}

	return courseStudent, nil
}

func (r *AuthStudentPostgres) UpdateCourseFileUrl(id int, file_url string) (course.Course, error) {
	var (
		courseByID course.Course
		url        string
	)
	query := fmt.Sprintf("SELECT id, name, description, student_group, file_url FROM %s WHERE id = $1", courseTable)

	row := r.db.QueryRow(query, id)
	err := row.Scan(&courseByID.ID, &courseByID.Name, &courseByID.Description, &courseByID.Student_group, &url)
	if err != nil {
		return courseByID, err
	}

	courseByID.File_url = make(map[int]string)
	err = json.Unmarshal([]byte(url), &courseByID.File_url)
	if err != nil {
		return courseByID, err
	}

	if courseByID.File_url == nil {
		courseByID.File_url = make(map[int]string)
	}

	courseByID.File_url[id] = file_url

	js, err := json.Marshal(courseByID.File_url)
	if err != nil {
		return courseByID, err
	}

	queryUpdate := fmt.Sprintf("UPDATE %s SET file_url = $1 WHERE id = $2", courseTable)
	_, err = r.db.Exec(queryUpdate, js, id)
	if err != nil {
		return courseByID, err
	}

	return courseByID, nil
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
