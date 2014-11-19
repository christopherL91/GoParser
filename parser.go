package main

import (
	"flag"
	"fmt"
	"github.com/christopherL91/Parser/toki"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

type Instruction struct {
	Name         string
	Num          int
	Color        string
	Instructions []Instruction
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
	defintions = []toki.Def{
		{Token: STRING, Pattern: `\"`},
		{Token: FORW, Pattern: `FORW\s+`},
		{Token: LEFT, Pattern: `LEFT\s+`},
		{Token: DOWN, Pattern: "DOWN"},
		{Token: BACK, Pattern: "BACK"},
		{Token: REP, Pattern: `REP\s+`},
		{Token: NUMBER, Pattern: "[0-9]+"},
		{Token: COLORKEYWORD, Pattern: `COLOR\s+`},
		{Token: RIGHT, Pattern: "RIGHT"},
		{Token: UP, Pattern: "UP"},
		{Token: COLOR, Pattern: `\#[A-F 0-9]{6}`},
		{Token: DOT, Pattern: `\.`},
		{Token: OTHER, Pattern: `.*`},
	}
	removeComments = regexp.MustCompile("%[^\\n]*")
	debug          bool
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
	fmt.Println()
}

func validateBuffer(buffer []*toki.Result) error {
	// instructions := []Instruction{}
L:
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
			// instruction := Instruction{
			// 	Name: "DOWN",
			// 	Num:  toInt(buffer[i+1].Value),
			// }
			// instruction = append(instructions, instruction)
			i += 2
		case FORW:
			// Next token must be a NUMBER
			if !(buffer[i+1].Token == NUMBER) {
				return syntaxError(buffer[i+1].Pos.Line)
			}
			if !(buffer[i+2].Token == DOT) {
				return syntaxError(buffer[i+2].Pos.Line)
			}
			// instruction := Instruction{
			// 	Name: "Forward",
			// 	Num:  toInt(buffer[i+1].Value),
			// }
			// instructions = append(instructions, instruction)
			i += 3
		case LEFT:
			if !(buffer[i+1].Token == NUMBER) {
				return syntaxError(buffer[i+1].Pos.Line)
			}
			if !(buffer[i+2].Token == DOT) {
				return syntaxError(buffer[i+2].Pos.Line)
			}
			i += 3
		case BACK:
			if !(buffer[i+1].Token == NUMBER) {
				return syntaxError(buffer[i+1].Pos.Line)
			}
			if !(buffer[i+2].Token == DOT) {
				return syntaxError(buffer[i+2].Pos.Line)
			}
			i += 3
		case RIGHT:
			if !(buffer[i+1].Token == NUMBER) {
				return syntaxError(buffer[i+1].Pos.Line)
			}
			if !(buffer[i+2].Token == DOT) {
				return syntaxError(buffer[i+2].Pos.Line)
			}
			i += 3
		case COLORKEYWORD:
			// Next token must be COLOR
			if !(buffer[i+1].Token == COLOR) {
				return syntaxError(buffer[i+1].Pos.Line)
			}
			//and after that DOT
			if !(buffer[i+2].Token == DOT) {
				return syntaxError(buffer[i+2].Pos.Line)
			}
			// instructions = append(instructions, []Instruction{ColorInstruction{red: 255,green:0,blue:255}})
			i += 3
		case REP:
			// Not implemented
			break L
			// 	if buffer[i+2].Token == STRING {
			// 		i += 3
			// 		for {

			// 			if buffer[i].Token == STRING {
			// 				instruction := Instruction{
			// 					Name:         "REP",
			// 					Instructions: parsed_instructions,
			// 				}
			// 				break
			// 			}
			// 		}
			// 	} else {
			// 		instruction := Instruction{
			// 			Name:         "REP",
			// 			Instructions: []Instruction{parsed_instruction},
			// 		}
			// 	}
			//
		default:
			// Found token that's invalid
			return syntaxError(buffer[i].Pos.Line)
		}
	}
	// Everything seems okey!
	return nil
}

func evaluateProgram(program []Instruction) string {
	// TODO
	return ""
}

func syntaxError(line int) error {
	return fmt.Errorf("Syntaxfel på rad %d", line)
}
