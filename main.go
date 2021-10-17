package main

import (
	"errors"
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
	name := "show_user"

	url, err := getURL(name)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return fail
	}

	if err := runAPI(url); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return fail
	}
	return success
}

func getURL(name string) (string, error) {
	buf, err := ioutil.ReadFile("./examples/requests.yml")
	if err != nil {
		return "", err
	}

	var reqs requests
	if err := yaml.Unmarshal(buf, &reqs); err != nil {
		return "", err
	}

	var req *request
	for _, r := range reqs {
		if r.Name == name {
			req = r
		}
	}
	if req == nil {
		return "", errors.New("request not found")
	}

	return req.URL, nil
}

func runAPI(url string) error {
	resp, err := http.Get(url)
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

type request struct {
	Name string
	URL  string
}

type requests []*request
