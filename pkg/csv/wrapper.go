package csv

import (
	"books-go/pkg/book"
	"books-go/pkg/logger"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Handle is a public function for accessing all functions in the csv package
func Handle() {
	r := mux.NewRouter()
	s := r.PathPrefix("/book").Subrouter()

	go logger.LogToFile()

	s.HandleFunc("/", wrapView).Methods("GET")
	s.HandleFunc("/{id}", wrapViewID).Methods("GET")
	s.HandleFunc("/", wrapInsert).Methods("POST")
	s.HandleFunc("/{id}", wrapUpdate).Methods("PUT")
	s.HandleFunc("/{id}", wrapDelete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8081", s))
}

// GET .../book/
func wrapView(w http.ResponseWriter, r *http.Request) {
	logger.Start()
	defer logger.End("task-1")
	defer logger.Log("task-1\t", "VIEW ALL")

	allBooks, err := view()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	fmt.Println(allBooks)
	json.NewEncoder(w).Encode(allBooks)
}

// GET .../book/{id}
func wrapViewID(w http.ResponseWriter, r *http.Request) {
	logger.Start()
	defer logger.End("task-1")

	vars := mux.Vars(r)
	key := vars["id"]

	id, _ := strconv.Atoi(key)

	op := fmt.Sprintf("VIEW id=%d", id)
	defer logger.Log("task-1\t", op)

	bookByID, err := viewID(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(bookByID)
}

// POST .../book/
func wrapInsert(w http.ResponseWriter, r *http.Request) {
	logger.Start()
	defer logger.End("task-1")
	reqBody, _ := ioutil.ReadAll(r.Body)

	var bk book.Book
	json.Unmarshal(reqBody, &bk)

	op := fmt.Sprintf("INSERT id=%d, revision=%d, isbn=%s, title=%s, author=%s", bk.ID, bk.Revision, bk.ISBN, bk.Title, bk.Author)
	defer logger.Log("task-1\t", op)

	err := insert(bk)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode("Success Insert")
}

// PUT .../book/{id}
func wrapUpdate(w http.ResponseWriter, r *http.Request) {
	logger.Start()
	defer logger.End("task-1")
	vars := mux.Vars(r)
	key := vars["id"]

	id, _ := strconv.Atoi(key)
	reqBody, _ := ioutil.ReadAll(r.Body)

	var bk book.Book
	json.Unmarshal(reqBody, &bk)

	op := fmt.Sprintf("UPDATE id=%d to id=%d, revision=%d, isbn=%s, title=%s, author=%s", id, bk.ID, bk.Revision, bk.ISBN, bk.Title, bk.Author)
	defer logger.Log("task-1\t", op)

	err := update(id, bk)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode("Success Update")
}

// DELETE .../book/{id}
func wrapDelete(w http.ResponseWriter, r *http.Request) {
	logger.Start()
	defer logger.End("task-1")
	vars := mux.Vars(r)
	key := vars["id"]

	id, _ := strconv.Atoi(key)

	op := fmt.Sprintf("DELETE id=%d", id)
	defer logger.Log("task-1\t", op)

	err := delete(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode("Success")
}
