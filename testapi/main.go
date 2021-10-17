package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/hokita/ac/testapi/handler"
)

const (
	success = 0
	fail    = 1
)

func main() {
	os.Exit(run())
}

func run() int {
	r := mux.NewRouter()
	r.Handle("/users/", &handler.IndexUsersHandler{}).Methods(http.MethodGet)
	r.Handle("/users/", &handler.CreateUserHandler{}).Methods(http.MethodPost)
	r.Handle("/users/{id:[0-9]+}/", &handler.ShowUserHandler{}).Methods(http.MethodGet)
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return fail
	}
	return success
}
