package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/go-yaml/yaml"
)

const (
	success = 0
	fail    = 1
)

func main() {
	os.Exit(run())
}

func run() int {
	flag.Parse()
	name := flag.Args()[0]

	path := os.Getenv("ACFILE")
	if path == "" {
		path = "./examples/requests.yml"
	}

	req, err := getRequest(path, name)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return fail
	}

	if err := runAPI(req); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return fail
	}
	return success
}

func getRequest(path, name string) (*request, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var reqs requests
	if err := yaml.Unmarshal(buf, &reqs); err != nil {
		return nil, err
	}

	var req *request
	for _, r := range reqs {
		if r.Name == name {
			req = r
		}
	}
	if req == nil {
		return nil, errors.New("request not found")
	}

	return req, nil
}

func runAPI(req *request) error {
	switch req.Method {
	case "get":
		resp, err := http.Get(req.URL)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		fmt.Println(string(body))

		return nil
	case "post":
		resp, err := http.Post(
			req.URL,
			"application/json",
			bytes.NewBuffer([]byte(req.JSON)),
		)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		fmt.Println(string(body))

		return nil
	}

	return errors.New("method not found")
}

type request struct {
	Name   string
	URL    string
	Method string
	JSON   string
}

type requests []*request
