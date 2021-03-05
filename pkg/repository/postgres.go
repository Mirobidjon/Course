package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	studentTable       = "students"
	masterTable        = "masters"
	courseTable        = "course"
	teacherCourseTable = "teacher_course"
	bookTable          = "books"
	studentBookTable   = "student_books"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
	SslMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			cfg.Host, cfg.Port, cfg.Username, cfg.DbName, cfg.Password, cfg.SslMode))

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, err
}
