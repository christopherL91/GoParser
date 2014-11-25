package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/christopherL91/Parser/toki"
)

// Used for evaluation
type Instruction struct {
	Name         string
	Num          int
	Color        Color
	Instructions []Instruction // Used within REP
}

// Defines a color using RGB color system.
type Color struct {
	Red, Green, Blue uint8
}

const (
	NUMBER toki.Token = iota + 1
	STRING
	DOWN
	FORW
	LEFT
	BACK
	RIGHT
	UP
	COLORKEYWORD
	REP
	DOT
	COLOR
	OTHER
)

var (
	debug          bool
	removeComments = regexp.MustCompile("%[^\\n]*")
	defintions     = []toki.Def{
		{Token: STRING, Pattern: `\"`},
		{Token: FORW, Pattern: `FORW\s+`},
		{Token: LEFT, Pattern: `LEFT\s+`},
		{Token: DOWN, Pattern: "DOWN"},
		{Token: BACK, Pattern: "BACK"},
		{Token: REP, Pattern: `REP\s+`},
		{Token: NUMBER, Pattern: `[0-9]+`},
		{Token: COLORKEYWORD, Pattern: `COLOR\s+`},
		{Token: RIGHT, Pattern: "RIGHT"},
		{Token: UP, Pattern: "UP"},
		{Token: COLOR, Pattern: `\#[A-F 0-9]{6}`},
		{Token: DOT, Pattern: `\.`},
		{Token: OTHER, Pattern: `.*`},
	}
)

func init() {
	flag.BoolVar(&debug, "debug", false, "Debuging")
	flag.Parse()
}

func main() {
	input, err := ioutil.ReadAll(os.Stdin) // Read input from stdin
	if err != nil {
		log.Fatal(err)
	}
	s := toki.NewScanner(defintions)
	noComments := removeComments.ReplaceAllString(string(input), "")
	s.SetInput(noComments)
	buffer := []*toki.Result{} // Holds all the tokens.
	for r := s.Next(); r.Token != toki.EOF; r = s.Next() {
		buffer = append(buffer, r) // Append new token to list.
	}
	if debug {
		prettyPrint(buffer)
		fmt.Println("Length of buffer:", len(buffer))
	}
	if err := validateBuffer(buffer); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Println("The program is valid... Beginning evaluation...")
}

func prettyPrint(buffer []*toki.Result) {
	defer fmt.Println()
	for _, val := range buffer {
		value := strings.TrimSpace(string(val.Value))
		switch val.Token {
		case NUMBER:
			fmt.Println("NUMBER", value)
		case DOWN:
			fmt.Println("DOWN", value)
		case FORW:
			fmt.Println("FORW", value)
		case LEFT:
			fmt.Println("LEFT", value)
		case BACK:
			fmt.Println("BACK", value)
		case RIGHT:
			fmt.Println("RIGHT", value)
		case UP:
			fmt.Println("UP", value)
		case COLORKEYWORD:
			fmt.Println("COLORKEYWORD", value)
		case REP:
			fmt.Println("REP", value)
		case DOT:
			fmt.Println("DOT", value)
		case STRING:
			fmt.Println("STRING", value)
		case COLOR:
			fmt.Println("COLOR", value)
		}
	}
}

func validateBuffer(buffer []*toki.Result) error {
	for i := 0; i < len(buffer); {
		switch buffer[i].Token {
		case UP:
			fallthrough
		case DOWN:
			if len(buffer) <= i+1 {
				return syntaxError(buffer[i].Pos.Line)
			}
			// Next token must be a DOT
			if !(buffer[i+1].Token == DOT) {
				return syntaxError(buffer[i+1].Pos.Line)
			}
			i += 2
		case FORW:
			if len(buffer) <= i+2 {
				return syntaxError(buffer[i].Pos.Line)
			}
			// Next token must be a NUMBER
			if !(buffer[i+1].Token == NUMBER) {
				return syntaxError(buffer[i+1].Pos.Line)
			}
			// Next token must be a DOT
			if !(buffer[i+2].Token == DOT) {
				return syntaxError(buffer[i+2].Pos.Line)
			}
			i += 3
		case LEFT:
			if len(buffer) <= i+2 {
				return syntaxError(buffer[i].Pos.Line)
			}
			// Next token must be a NUMBER
			if !(buffer[i+1].Token == NUMBER) {
				return syntaxError(buffer[i+1].Pos.Line)
			}
			// Next token must be a DOT
			if !(buffer[i+2].Token == DOT) {
				return syntaxError(buffer[i+2].Pos.Line)
			}
			i += 3
		case BACK:
			if len(buffer) <= i+2 {
				return syntaxError(buffer[i].Pos.Line)
			}
			// Next token must be a NUMBER
			if !(buffer[i+1].Token == NUMBER) {
				return syntaxError(buffer[i+1].Pos.Line)
			}
			// Next token must be a DOT
			if !(buffer[i+2].Token == DOT) {
				return syntaxError(buffer[i+2].Pos.Line)
			}
			i += 3
		case RIGHT:
			if len(buffer) <= i+2 {
				return syntaxError(buffer[i].Pos.Line)
			}
			// Next token must be a NUMBER
			if !(buffer[i+1].Token == NUMBER) {
				return syntaxError(buffer[i+1].Pos.Line)
			}
			// Next token must be a DOT
			if !(buffer[i+2].Token == DOT) {
				return syntaxError(buffer[i+2].Pos.Line)
			}
			i += 3
		case COLORKEYWORD:
			if len(buffer) <= i+2 {
				return syntaxError(buffer[i].Pos.Line)
			}
			// Next token must be COLOR
			if !(buffer[i+1].Token == COLOR) {
				return syntaxError(buffer[i+1].Pos.Line)
			}
			//and after that DOT
			if !(buffer[i+2].Token == DOT) {
				return syntaxError(buffer[i+2].Pos.Line)
			}
			i += 3
		case REP:
			if len(buffer) <= i+3 {
				return syntaxError(buffer[i].Pos.Line)
			}
			if !(buffer[i+1].Token == NUMBER) {
				return syntaxError(buffer[i+1].Pos.Line)
			}
			if !(buffer[i+2].Token == STRING) {
				return syntaxError(buffer[i+2].Pos.Line)
			}
			buf := bytes.NewBuffer(buffer[i+1].Value) // buffer[i+1].Value is []byte
			repetitions, err := binary.ReadVarint(buf)
			if err != nil {
				// This should not happen, because of regexp check.
				return syntaxError(buffer[i+1].Pos.Line)
			}
			/*
					repetitions
				------------------------------
				REP NUMBER     |   STRING .... STRING
				NOW	CHECKED    |   CHECKED

				* Move 3 tokens
			*/

			i += 3
			for {
				if buffer[i].Token == STRING {
					fmt.Println("String is now complete")
					break
				}
			}
		default:
			// Found token that's invalid
			return syntaxError(buffer[i].Pos.Line)
		}
	}
	return nil // Validation complete. Everything seems to work!
}

func evaluateProgram(program []Instruction) string {
	// TODO
	return ""
}

func syntaxError(line int) error {
	return fmt.Errorf("Syntaxfel på rad %d", line)
}
