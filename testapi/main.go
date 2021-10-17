package main

import (
	"fmt"
	"net/http"
	"os"

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
	http.Handle("/users/", &handler.IndexHandler{})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return fail
	}
	return success
}
