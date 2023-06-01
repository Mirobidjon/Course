package service

import (
	"github.com/Mirobidjon/course"
	"github.com/Mirobidjon/course/pkg/repository"
)

type AuthStudents interface {
	CreateStudent(input course.SignUpStudent) (int, error)
	GenerateToken(username, password string) (string, error)
	GetAllCourse(id int) ([]course.Course, error)
	GetCourseByID(courseID, studentID int) (course.Course, error)
	GetAllStudents() ([]course.GetStudents, error)
	GetStudentByID(id int) (course.GetStudents, error)
	UpdateStudent(input course.UpdateStudent, id int) error
	DeleteStudent(id int) error
	UpdateCourseFileUrl(id int, file_url string) (course.Course, error)
}

type AuthMasters interface {
	CreateMaster(input course.MasterInput, role string) (int, error)
	GenerateTokenMaster(username, password, role string) (string, error)
	ParseToken(tk string) (string, int, error)
	GetAllTeacherCourse() ([]course.Course, error)
	GetAllMaster(role string) ([]course.OutputMaster, error)
	GetMasterByID(role string, id int) (course.OutputMaster, error)
	UpdateMaster(input course.MasterInput, id int) error
	DeleteTeacher(role string, id int) error
}

type Course interface {
	CreateCourse(input course.InputCourse, teacherID int) (int, error)
	DeleteCourse(courseID, teacherID int) error
	UpdateCourse(name, description string, courseID, teacherID int) error
	GetAllCourse(teacherID int) ([]course.Course, error)
	GetTeacherCourse(teacherID, courseID int) (course.Course, error)
	GetCourse(courseID int) (course.Course, error)
}

type Book interface {
	CreateBook(name, author string, studentID int) (int, error)
	GetBookByID(bookID, studentID int) (course.Book, error)
	GetAllBooks(studentID int) ([]course.Book, error)
	DeleteBook(studentID, bookID int) error
	UpdateBook(name, author string, studentID, bookID int) error
}

type Service struct {
	AuthStudents
	AuthMasters
	Course
	Book
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		AuthStudents: NewAuthStudentService(repo.AuthStudents),
		AuthMasters:  NewAuthMasterService(repo.AuthMasters),
		Book:         NewBookService(repo.Book),
		Course:       NewCourseService(repo.Course),
	}
}
