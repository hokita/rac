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
	r.Handle("/users/", &handler.IndexHandler{}).Methods(http.MethodGet)
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return fail
	}
	return success
}
