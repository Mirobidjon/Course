package service

import (
	"errors"
	"github.com/Mirobidjon/course"
	"github.com/Mirobidjon/course/pkg/repository"
)

type AuthMasterService struct {
	repo repository.AuthMasters
}

func NewAuthMasterService(repo repository.AuthMasters) *AuthMasterService {
	return &AuthMasterService{repo: repo}
}

func (r *AuthMasterService) CreateMaster(input course.MasterInput, role string) (int, error) {
	input.Password = generatePassword(input.Password)
	return r.repo.CreateMaster(input, role)
}

func (r *AuthMasterService) GenerateTokenMaster(username, password, role string) (string, error) {
	id, err := r.repo.GetMaster(username, generatePassword(password), role)
	if err != nil {
		return "", err
	}

	return GetToken(id, role)
}

func (r *AuthMasterService) GetAllTeacherCourse() ([]course.Course, error) {
	return r.repo.GetAllTeacherCourse()
}

func (r *AuthMasterService) GetAllMaster(role string) ([]course.OutputMaster, error) {
	return r.repo.GetAllMaster(role)
}

func (r *AuthMasterService) GetMasterByID(role string, id int) (course.OutputMaster, error) {
	return r.repo.GetMasterByID(role, id)
}

func (r *AuthMasterService) UpdateMaster(input course.MasterInput, id int) error {
	if input.Name == "" && input.Username == "" && input.Password == "" {
		return errors.New("Update hasn't values. ")
	}

	if input.Password != "" {
		input.Password = generatePassword(input.Password)
	}

	return r.repo.UpdateMaster(input, id)
}

func (r *AuthMasterService) DeleteTeacher(role string, id int) error {
	return r.repo.DeleteTeacher(role, id)
}
