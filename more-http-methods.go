package main

import (
	"net/http"

	"log"
)

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{message: get working}`))
	case "POST":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{message: post working}`))
	case "PUT":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{message: put working}`))
	default:
		w.WriteHeader(http.StatusNotFound)
        w.Write([]byte(`{"message": "not found"}`))
	}

}

func main() {
	s:= &server{}
	http.Handle("/", s)
	log.Fatal(http.ListenAndServe(":8000", nil))

}