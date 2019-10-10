package main

import (
	"io/ioutil"
	"strings"
)

var adjectives []string
var nouns []string

// Initialize ...
func Initialize() error {
	fileContent, err := ioutil.ReadFile("./words/adjectives.txt")
	if err != nil {
		return err
	}
	adjectives = strings.Split(string(fileContent), " ")
	fileContent, err = ioutil.ReadFile("./words/nouns.txt")
	if err != nil {
		return err
	}
	nouns = strings.Split(string(fileContent), " ")
	return nil
}

func main() {
	Initialize()
}
