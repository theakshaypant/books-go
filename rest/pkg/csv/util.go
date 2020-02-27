package csv

import (
	"books-go/pkg/book"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/rogpeppe/go-internal/lockedfile"
)

var datafile *lockedfile.File

const (
	datapath = "../../tools/data.csv"
)

func open(flag int) error {
	var err error
	datafile, err = lockedfile.OpenFile(datapath, flag, 0644)
	if err != nil {
		return err
	}
	return nil
}

func view() ([]book.Book, error) {
	err := open(os.O_RDONLY)
	if err != nil {
		return []book.Book{}, err
	}

	data, err := csv.NewReader(datafile).ReadAll()
	if err != nil {
		return []book.Book{}, err
	}
	close()

	var books []book.Book
	for _, record := range data {
		idTemp, _ := strconv.Atoi(record[0])
		revTemp, _ := strconv.Atoi(record[1])
		bookTemp := book.Book{
			ID:       idTemp,
			Revision: revTemp,
			ISBN:     record[2],
			Title:    record[3],
			Author:   record[4],
		}
		books = append(books, bookTemp)
	}
	return books[:], nil
}

func viewID(id int) (book.Book, error) {
	data, err := view()
	if err != nil {
		return book.Book{}, err
	}

	for _, record := range data {
		if record.ID == id {
			return record, nil
		}
	}
	return book.Book{}, errors.New("id not found")
}

func insert(bk book.Book) error {
	_, err := viewID(bk.ID)
	if err == nil {
		return errors.New("id exists") // ID already exists
	}

	data, err := view()
	if err != nil {
		return err
	}

	err = open(os.O_WRONLY)
	datafile.Truncate(0)
	datafile.Seek(0, 0)
	wo := csv.NewWriter(datafile)
	for _, record := range data {
		temp := []string{strconv.Itoa(record.ID), strconv.Itoa(record.Revision), record.ISBN, record.Title, record.Author}
		if err = wo.Write(temp); err != nil {
			return err
		}
	}
	temp := []string{strconv.Itoa(bk.ID), strconv.Itoa(bk.Revision), bk.ISBN, bk.Title, bk.Author}
	if err = wo.Write(temp); err != nil {
		return err
	}
	wo.Flush()
	close()

	return nil
}

func update(id int, bk book.Book) error {
	data, err := view()
	if err != nil {
		return err
	}

	err = open(os.O_WRONLY)
	datafile.Truncate(0)
	datafile.Seek(0, 0)
	wo := csv.NewWriter(datafile)
	for _, record := range data {
		if record.ID != id {
			temp := []string{strconv.Itoa(record.ID), strconv.Itoa(record.Revision), record.ISBN, record.Title, record.Author}
			if err = wo.Write(temp); err != nil {
				return err
			}
		} else {
			temp := []string{strconv.Itoa(bk.ID), strconv.Itoa(bk.Revision), bk.ISBN, bk.Title, bk.Author}
			fmt.Println(temp)
			if err = wo.Write(temp); err != nil {
				return err
			}
		}
	}
	wo.Flush()
	close()

	return nil
}

func delete(id int) error {
	data, err := view()
	if err != nil {
		return err
	}

	err = open(os.O_WRONLY)
	datafile.Truncate(0)
	datafile.Seek(0, 0)

	wo := csv.NewWriter(datafile)
	for _, record := range data {
		if record.ID != id {
			temp := []string{strconv.Itoa(record.ID), strconv.Itoa(record.Revision), record.ISBN, record.Title, record.Author}
			if err = wo.Write(temp); err != nil {
				return err
			}
		}
	}
	wo.Flush()
	close()

	return nil
}

func close() {
	datafile.Close()
}
