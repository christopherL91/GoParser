package main

import (
	"fmt"
	"github.com/christopherL91/toki"
	"io/ioutil"
	"log"
	"os"
	"unsafe"
)

const (
	NUMBER toki.Token = iota + 1
	DOWN
	FORW
	LEFT
	COMMENT
	BACK
	RIGHT
	UP
	COLOR
	REP
	WHITESPACE
	DOT
)

var (
	defintions = []toki.Def{
		{Token: NUMBER, Pattern: "[0-9]+"},
		{Token: COMMENT, Pattern: "%.*\n"},
		{Token: FORW, Pattern: "FORW"},
		{Token: LEFT, Pattern: "LEFT"},
		{Token: DOWN, Pattern: "DOWN"},
		{Token: WHITESPACE, Pattern: `(\s)+`},
		{Token: BACK, Pattern: "BACK"},
		{Token: REP, Pattern: "REP"},
		{Token: COLOR, Pattern: "COLOR"},
		{Token: RIGHT, Pattern: "RIGHT"},
		{Token: UP, Pattern: "UP"},
		{Token: DOT, Pattern: "."},
	}
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin) // Read input from stdin
	if err != nil {
		log.Fatal(err)
	}
	s := toki.NewScanner(defintions)
	s.SetInput(string(input))
	buffer := []*toki.Result{} // Holds all the tokens.
	for r := s.Next(); r.Token != toki.EOF; r = s.Next() {
		buffer = append(buffer, r) // Append new token to list.
	}
	validateBuffer(buffer)
}

func validateBuffer(buffer []*toki.Result) {
	fmt.Println(buffer)
	fmt.Println(unsafe.Sizeof(buffer))
}
