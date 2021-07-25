package main

import (
	"net/http"

	"log"
)

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{message: its working}`))
}

func main() {
	s:= &server{}
	http.Handle("/", s)
	log.Fatal(http.ListenAndServe(":8000", nil))

}