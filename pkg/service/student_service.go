package service

import (
	"errors"

	"github.com/Mirobidjon/course"
	"github.com/Mirobidjon/course/pkg/repository"
)

type AuthStudentService struct {
	repo repository.AuthStudents
}

func NewAuthStudentService(repo repository.AuthStudents) *AuthStudentService {
	return &AuthStudentService{repo: repo}
}

func (s AuthStudentService) CreateStudent(input course.SignUpStudent) (int, error) {
	input.Password = generatePassword(input.Password)
	return s.repo.CreateStudent(input)
}

func (s AuthStudentService) GenerateToken(username, password string) (string, error) {
	studentID, err := s.repo.GetStudent(username, generatePassword(password))
	if err != nil {
		return "", err
	}

	return GetToken(studentID, "student")
}

func (s AuthStudentService) GetAllCourse(id int) ([]course.Course, error) {
	return s.repo.GetAllCourse(id)
}

func (s *AuthStudentService) GetAllStudents() ([]course.GetStudents, error) {
	return s.repo.GetAllStudents()
}

func (s *AuthStudentService) GetStudentByID(id int) (course.GetStudents, error) {
	return s.repo.GetStudentByID(id)
}

func (s *AuthStudentService) UpdateStudent(input course.UpdateStudent, id int) error {
	if input.Name == "" && input.Username == "" && input.Password == "" {
		return errors.New("Update hasn't values ")
	}

	if input.Password != "" {
		input.Password = generatePassword(input.Password)
	}

	return s.repo.UpdateStudent(input, id)
}

func (s *AuthStudentService) GetCourseByID(courseID, studentID int) (course.Course, error) {
	return s.repo.GetCourseByID(courseID, studentID)
}

func (s *AuthStudentService) DeleteStudent(id int) error {
	return s.repo.DeleteStudent(id)
}

func (s *AuthStudentService) UpdateCourseFileUrl(id int, file_url string) (course.Course, error) {
	return s.repo.UpdateCourseFileUrl(id, file_url)
}
