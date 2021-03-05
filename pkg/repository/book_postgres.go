package repository

import (
	"fmt"
	"github.com/Mirobidjon/course"
	"github.com/jmoiron/sqlx"
	"strings"
)

type BookPostgres struct {
	db *sqlx.DB
}

func NewBookPostgres(db *sqlx.DB) *BookPostgres {
	return &BookPostgres{db: db}
}

func (r *BookPostgres) CreateBook(name, author string, studentID int) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	queryBook := fmt.Sprintf("INSERT INTO %s (name, author) VALUES ($1, $2) RETURNING id",
		bookTable)
	row := tx.QueryRow(queryBook, name, author)
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}

	queryStudentBook := fmt.Sprintf("INSERT INTO %s (student_id, book_id) VALUES($1, $2)",
		studentBookTable)
	_, err = tx.Exec(queryStudentBook, studentID, id)
	if err != nil {
		return 0, err
	}

	return id, tx.Commit()
}

func (r *BookPostgres) GetBookByID(bookID, studentID int) (course.Book, error) {
	var bk course.Book
	query :=
		fmt.Sprintf("SELECT bl.id, bl.name, bl.author FROM %s bl INNER JOIN %s sl on bl.id = sl.book_id WHERE sl.student_id = $1 AND sl.book_id = $2",
			bookTable, studentBookTable)

	err := r.db.Get(&bk, query, studentID, bookID)

	return bk, err
}

func (r *BookPostgres) GetAllBooks(studentID int) ([]course.Book, error) {
	var books []course.Book

	query := fmt.Sprintf(
		`SELECT bl.id, bl.name, bl.author FROM %s bl INNER JOIN %s sl on sl.book_id = bl.id 
				WHERE sl.student_id = $1`,
		bookTable, studentBookTable)

	if err := r.db.Select(&books, query, studentID); err != nil {
		return nil, err
	}

	return books, nil
}

func (r *BookPostgres) DeleteBook(studentID, bookID int) error {
	query := fmt.Sprintf(
		`DELETE FROM %s bl USING %s sl 
			WHERE sl.book_id = bl.id AND sl.student_id = $1 AND sl.book_id = $2`,
		bookTable, studentBookTable)

	_, err := r.db.Exec(query, studentID, bookID)

	return err
}

func (r *BookPostgres) UpdateBook(name, author string, studentID, bookID int) error {
	values := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if name != "" {
		values = append(values, fmt.Sprintf("name=$%d", argID))
		args = append(args, name)
		argID++
	}

	if author != "" {
		values = append(values, fmt.Sprintf("author=$%d", argID))
		args = append(args, author)
		argID++
	}

	setQuery := strings.Join(values, ", ")
	query :=
		fmt.Sprintf("UPDATE %s bl SET %s FROM %s sl WHERE bl.id = sl.book_id AND sl.student_id = $%d AND sl.book_id = $%d ",
			bookTable, setQuery, studentBookTable, argID, argID+1)

	args = append(args, studentID, bookID)

	_, err := r.db.Exec(query, args...)
	return err
}
