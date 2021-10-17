package main

import (
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
	url, err := getFile()
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

func getFile() (string, error) {
	buf, err := ioutil.ReadFile("./examples/api.yml")
	if err != nil {
		return "", err
	}

	var a apis
	if err := yaml.Unmarshal(buf, &a); err != nil {
		return "", err
	}

	return a[0].URL, nil
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

type api struct {
	Name   string
	URL    string
	Method string
}

type apis []*api
