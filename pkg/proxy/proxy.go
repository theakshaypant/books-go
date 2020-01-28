package proxy

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

const (
	task1 = "http://task1:8081/"
	task2 = "http://task2:8082/"
)

// Handle is a public function which redirects requests to the appropriate package server in turn
func Handle() {
	r := mux.NewRouter()
	r.HandleFunc("/v1/book/", csvHandler)
	r.HandleFunc("/v1/book/{id}", csvHandler)
	r.HandleFunc("/v2/book/", dbHandler)
	r.HandleFunc("/v2/book/{id}", dbHandler)
	http.ListenAndServe(":8080", r)
}

func csvHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/v1/")
	//fmt.Println(path, r.Method, r.Body)
	client := &http.Client{}
	url := fmt.Sprintf("%s%s", task1, path)
	//	fmt.Println(url)
	req, err := http.NewRequest(r.Method, url, r.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	data, err := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(data), err)
	w.Write(data)
}

func dbHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/v2/")
	// fmt.Println(path, r.Method, r.Body)
	client := &http.Client{}
	url := fmt.Sprintf("%s%s", task2, path)
	// fmt.Println(url)
	req, err := http.NewRequest(r.Method, url, r.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	data, err := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(data), err)
	w.Write(data)
}
