package service

import (
	"errors"
	"github.com/Mirobidjon/course"
	"github.com/Mirobidjon/course/pkg/repository"
)

type CourseService struct {
	repo repository.Course
}

func NewCourseService(repo repository.Course) *CourseService {
	return &CourseService{repo: repo}
}

func (s *CourseService) CreateCourse(input course.InputCourse, teacherID int) (int, error) {
	if input.Name == "" || input.Group <= 0 {
		return 0, errors.New("Invalid course name or group! ")
	}
	return s.repo.CreateCourse(input, teacherID)
}

func (s *CourseService) DeleteCourse(courseID, teacherID int) error {
	return s.repo.DeleteCourse(courseID, teacherID)
}

func (s *CourseService) UpdateCourse(name, description string, courseID, teacherID int) error {
	if name == "" && description != "" {
		return errors.New("Update course hasn't value. ")
	}
	return s.repo.UpdateCourse(name, description, courseID, teacherID)
}

func (s *CourseService) GetAllCourse(teacherID int) ([]course.Course, error) {
	return s.repo.GetAllCourse(teacherID)
}

func (s *CourseService) GetTeacherCourse(teacherID, courseID int) (course.Course, error) {
	return s.repo.GetTeacherCourse(teacherID, courseID)
}

func (s *CourseService) GetCourse(courseID int) (course.Course, error) {
	return s.repo.GetCourse(courseID)
}
