package models

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB

type Book struct {
	Title  string
	Author string
	Read   string
}

// Get all books in the books table.
func AllBooks() ([]Book, error) {
	query := "SELECT * FROM books"
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bks, err := makeBookSlice(rows)
	if err != nil {
		return nil, err
	}

	return bks, nil
}

// Query for books by name. 
func NameQuery(r string) ([]Book, error) {
	rows, err := DB.Query("SELECT * FROM books WHERE name = ?", r)
	
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bks, err := makeBookSlice(rows)
	if err != nil {
		return nil, err
	}

	return bks, nil
}

// Query for books by Author. 
func AuthorQuery(r string) ([]Book, error) {
	rows, err := DB.Query("SELECT * FROM books WHERE author = ?", r)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bks, err := makeBookSlice(rows)
	if err != nil {
		return nil, err
	}

	return bks, nil
}

// Query for books by read.  
func ReadQuery(r string) ([]Book, error) {
	rows, err := DB.Query("SELECT * FROM books WHERE read = ?", r)
	
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bks, err := makeBookSlice(rows)
	if err != nil {
		return nil, err
	}

	return bks, nil
}

// A helper function to cast the query results to a slice
func makeBookSlice(r *sql.Rows) ([]Book, error) {
	var bks []Book

	for r.Next() {
		var bk Book

		err := r.Scan(&bk.Title, &bk.Author, &bk.Read)
		if err != nil {
			return nil, err
		}

		bks = append(bks, bk)
	}
	if err := r.Err(); err != nil {
		return nil, err
	}

	return bks, nil
}
