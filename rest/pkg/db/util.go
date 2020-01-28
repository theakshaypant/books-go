package db

import (
	"books-go/pkg/book"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // helper import to access postgres db
)

var db *sql.DB

const (
	host     = "db"
	port     = 5432
	user     = "postgres"
	password = "thanks.123"
	dbname   = "books_db"
)

func open() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	// change query
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS books_db (id INT UNIQUE, revision INT, isbn TEXT, title TEXT, author TEXT);")
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func view() ([]book.Book, error) {
	err := open()
	if err != nil {
		return []book.Book{}, err
	}
	var books []book.Book

	sqlStatement := fmt.Sprintf("SELECT * FROM %s;", dbname)
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return []book.Book{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var bk book.Book
		err := rows.Scan(&bk.ID, &bk.Revision, &bk.ISBN,
			&bk.Title, &bk.Author)
		if err != nil {
			return []book.Book{}, err
		}

		books = append(books[:], bk)
	}
	close()
	return books, nil
}

func viewID(id int) (book.Book, error) {
	err := open()
	if err != nil {
		return book.Book{}, err
	}

	sqlStatement := `SELECT * FROM books_db WHERE id=$1;`
	var bk book.Book
	row := db.QueryRow(sqlStatement, id)
	switch err := row.Scan(&bk.ID, &bk.Revision, &bk.ISBN,
		&bk.Title, &bk.Author); err {
	case sql.ErrNoRows:
		return book.Book{}, err //errors.New("id does not exist")
	case nil:
		return bk, nil
	default:
		return book.Book{}, err
	}
}

func insert(bk book.Book) error {
	err := open()
	if err != nil {
		return err
	}
	defer close()

	sqlStatement := `INSERT INTO books_db (id, revision, isbn, title, author) VALUES ($1, $2, $3, $4, $5);`
	_, err = db.Exec(sqlStatement, bk.ID, bk.Revision, bk.ISBN, bk.Title, bk.Author)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func update(id int, bk book.Book) error {
	err := open()
	if err != nil {
		return err
	}

	sqlStatement := `UPDATE books_db SET revision=$1, isbn=$2, title=$3, author=$4, id=$5 WHERE id=$6;`
	_, err = db.Exec(sqlStatement, bk.Revision, bk.ISBN, bk.Title, bk.Author, bk.ID, id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	close()
	return nil
}

func delete(id int) error {
	err := open()
	if err != nil {
		return err
	}
	defer close()

	sqlStatement := `DELETE FROM books_db WHERE id=$1;`
	_, err = db.Exec(sqlStatement, id)
	if err != nil {
		return err
	}

	return nil
}

func close() {
	db.Close()
}
