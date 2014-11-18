package main

import (
	"fmt"
	"github.com/christopherL91/Parser/toki"
	"io/ioutil"
	"log"
	"os"
	"unsafe"
)

const (
	NUMBER toki.Token = iota + 1
	STRING
	DOWN
	FORW
	LEFT
	COMMENT
	BACK
	RIGHT
	UP
	COLORKEYWORD
	REP
	WHITESPACE
	DOT
	COLOR
)

var (
	defintions = []toki.Def{
		{Token: COLOR, Pattern: `\#[A-F 0-9]{6}`},
		{Token: STRING, Pattern: `\"`},
		{Token: COMMENT, Pattern: "%.*\n"},
		{Token: FORW, Pattern: `FORW\s+`},
		{Token: LEFT, Pattern: `LEFT\s+`},
		{Token: DOWN, Pattern: "DOWN"},
		{Token: BACK, Pattern: "BACK"},
		{Token: REP, Pattern: `REP\s+`},
		{Token: NUMBER, Pattern: "[0-9]+"},
		{Token: COLORKEYWORD, Pattern: `COLOR\s+`},
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
	prettyPrint(buffer)
	// validateBuffer(buffer)
}

func prettyPrint(buffer []*toki.Result) {
	for _, val := range buffer {
		switch val.Token {
		case NUMBER:
			fmt.Println("NUMBER")
		case DOWN:
			fmt.Println("DOWN")
		case FORW:
			fmt.Println("FORW")
		case LEFT:
			fmt.Println("LEFT")
		case COMMENT:
			fmt.Println("COMMENT")
		case BACK:
			fmt.Println("BACK")
		case RIGHT:
			fmt.Println("RIGHT")
		case UP:
			fmt.Println("UP")
		case COLORKEYWORD:
			fmt.Println("COLORKEYWORD")
		case REP:
			fmt.Println("REP")
		case WHITESPACE:
			fmt.Println("WHITESPACE")
		case DOT:
			fmt.Println("DOT")
		case STRING:
			fmt.Println("STRING")
		case COLOR:
			fmt.Println("COLOR")
		}
	}
}

func validateBuffer(buffer []*toki.Result) {
	for index, token := range buffer {
		switch token.Token {
		case NUMBER:

		case DOWN:

		case FORW:

		case LEFT:

		case COMMENT:
			// Do nothing
		case BACK:

		case RIGHT:

		case UP:

		case COLORKEYWORD:

		case REP:

		case WHITESPACE:

		case DOT:

		case STRING:

		case COLOR:
		}
	}
}
