package service

import (
	"errors"
	"github.com/Mirobidjon/course"
	"github.com/Mirobidjon/course/pkg/repository"
)

type BookService struct {
	repo repository.Book
}

func NewBookService(repo repository.Book) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) CreateBook(name, author string, studentID int) (int, error) {
	if name == "" || author == "" {
		return 0, errors.New("Invalid book name or book author !")
	}
	return s.repo.CreateBook(name, author, studentID)
}

func (s *BookService) GetBookByID(bookID, studentID int) (course.Book, error) {
	return s.repo.GetBookByID(bookID, studentID)
}

func (s *BookService) GetAllBooks(studentID int) ([]course.Book, error) {
	return s.repo.GetAllBooks(studentID)
}

func (s *BookService) DeleteBook(studentID, bookID int) error {
	return s.repo.DeleteBook(studentID, bookID)
}

func (s *BookService) UpdateBook(name, author string, studentID, bookID int) error {
	if name == "" && author == "" {
		return errors.New("update book hasn't value")
	}
	return s.repo.UpdateBook(name, author, studentID, bookID)
}
