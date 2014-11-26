package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/christopherL91/Parser/toki"
)

type instruction struct {
	name         string
	val          int
	color        string
	instructions []instruction
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
	flag.BoolVar(&debug, "debug", false, "Debug")
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
	inst, _, err := parse(buffer, 0, false, false)
	if err != nil {
		log.Fatalln(err)
		os.Exit(-1)
	}
	var buf bytes.Buffer
	pen := newPen()
	evaluateProgram(inst, &buf, pen)
	fmt.Print(buf.String())
}

func parse(buffer []*toki.Result, i int, parseone bool, inrep bool) ([]instruction, int, error) {
	instructions := []instruction{}
	for i < len(buffer) {
		// fmt.Println(i)
		switch buffer[i].Token {
		case STRING:
			if inrep {
				return instructions, i, nil
			} else {
				return nil, 0, syntaxError(buffer[i].Pos.Line)
			}
		case UP:
			inst := instruction{
				name: "UP",
			}
			instructions = append(instructions, inst)
			i += 2
		case DOWN:
			if len(buffer) <= i+1 {
				return nil, 0, syntaxError(buffer[i].Pos.Line)
			}
			// Next token must be a DOT
			if !(buffer[i+1].Token == DOT) {
				return nil, 0, syntaxError(buffer[i+1].Pos.Line)
			}
			inst := instruction{
				name: "DOWN",
			}
			instructions = append(instructions, inst)
			i += 2
		case FORW:
			if len(buffer) <= i+2 {
				return nil, 0, syntaxError(buffer[i].Pos.Line)
			}
			// Next token must be a NUMBER
			if !(buffer[i+1].Token == NUMBER) {
				return nil, 0, syntaxError(buffer[i+1].Pos.Line)
			}
			// Next token must be a DOT
			if !(buffer[i+2].Token == DOT) {
				return nil, 0, syntaxError(buffer[i+2].Pos.Line)
			}
			num, _ := strconv.Atoi(string(buffer[i+1].Value))
			inst := instruction{
				name: "FORW",
				val:  num,
			}
			instructions = append(instructions, inst)
			i += 3
		case LEFT:
			if len(buffer) <= i+2 {
				return nil, 0, syntaxError(buffer[i].Pos.Line)
			}
			// Next token must be a NUMBER
			if !(buffer[i+1].Token == NUMBER) {
				return nil, 0, syntaxError(buffer[i+1].Pos.Line)
			}
			// Next token must be a DOT
			if !(buffer[i+2].Token == DOT) {
				return nil, 0, syntaxError(buffer[i+2].Pos.Line)
			}
			num, _ := strconv.Atoi(string(buffer[i+1].Value))
			inst := instruction{
				name: "LEFT",
				val:  num,
			}
			instructions = append(instructions, inst)
			i += 3
		case BACK:
			if len(buffer) <= i+2 {
				return nil, 0, syntaxError(buffer[i].Pos.Line)
			}
			// Next token must be a NUMBER
			if !(buffer[i+1].Token == NUMBER) {
				return nil, 0, syntaxError(buffer[i+1].Pos.Line)
			}
			// Next token must be a DOT
			if !(buffer[i+2].Token == DOT) {
				return nil, 0, syntaxError(buffer[i+2].Pos.Line)
			}
			num, _ := strconv.Atoi(string(buffer[i+1].Value))
			inst := instruction{
				name: "BACK",
				val:  num,
			}
			instructions = append(instructions, inst)
			i += 3
		case RIGHT:
			if len(buffer) <= i+2 {
				return nil, 0, syntaxError(buffer[i].Pos.Line)
			}
			// Next token must be a NUMBER
			if !(buffer[i+1].Token == NUMBER) {
				return nil, 0, syntaxError(buffer[i+1].Pos.Line)
			}
			// Next token must be a DOT
			if !(buffer[i+2].Token == DOT) {
				return nil, 0, syntaxError(buffer[i+2].Pos.Line)
			}
			num, _ := strconv.Atoi(string(buffer[i+1].Value))
			inst := instruction{
				name: "RIGHT",
				val:  num,
			}
			instructions = append(instructions, inst)
			i += 3
		case COLORKEYWORD:
			if len(buffer) <= i+2 {
				return nil, 0, syntaxError(buffer[i].Pos.Line)
			}
			// Next token must be COLOR
			if !(buffer[i+1].Token == COLOR) {
				return nil, 0, syntaxError(buffer[i+1].Pos.Line)
			}
			//and after that DOT
			if !(buffer[i+2].Token == DOT) {
				return nil, 0, syntaxError(buffer[i+2].Pos.Line)
			}
			inst := instruction{
				name:  "COLOR",
				color: strings.ToUpper(string(buffer[i+1].Value)),
			}
			instructions = append(instructions, inst)
			i += 3
		case REP:
			if len(buffer) <= i+3 {
				return nil, 0, syntaxError(buffer[i].Pos.Line)
			}
			if !(buffer[i+1].Token == NUMBER) {
				return nil, 0, syntaxError(buffer[i+1].Pos.Line)
			}
			if buffer[i+2].Token == STRING {
				subinstructions, nextpos, err := parse(buffer, i+3, false, true)
				if err != nil {
					return nil, 0, err
				}
				numberLine := buffer[i+1].Pos.Line
				stringLine := buffer[i+2].Pos.Line
				numberCol := buffer[i+1].Pos.Column
				stringCol := buffer[i+2].Pos.Column
				// They are on the same line
				if numberLine == stringLine {
					// 5" <- illegal
					if stringCol-numberCol == 1 {
						return nil, 0, syntaxError(buffer[i+1].Pos.Line)
					}
				}
				num, _ := strconv.Atoi(string(buffer[i+1].Value))
				inst := instruction{
					name:         "REP",
					val:          num,
					instructions: subinstructions,
				}
				// No tokens left, and no `"` has shown up before.
				if len(buffer[nextpos:]) < 1 {
					return nil, 0, syntaxError(buffer[nextpos-1].Pos.Line)
				}
				if !(buffer[nextpos].Token == STRING) {
					return nil, 0, syntaxError(buffer[nextpos].Pos.Line)
				}
				instructions = append(instructions, inst)
				i = nextpos + 1
			} else {
				subinstructions, nextpos, err := parse(buffer, i+2, true, false)
				if err != nil {
					return nil, 0, err
				}
				num, _ := strconv.Atoi(string(buffer[i+1].Value))
				inst := instruction{
					name:         "REP",
					val:          num,
					instructions: subinstructions,
				}
				instructions = append(instructions, inst)
				i = nextpos
			}
		default:
			// Found token that's invalid
			return nil, 0, syntaxError(buffer[i].Pos.Line)
		}
		if parseone {
			return instructions, i, nil
		}
	}
	return instructions, i, nil // Validation complete. Everything seems to work!
}

func evaluateProgram(program []instruction, buffer *bytes.Buffer, pen *Pen) string {
	for _, inst := range program {
		switch inst.name {
		case "FORW":
			if pen.down {
				buffer.WriteString(pen.color)
				buffer.WriteRune(' ')
				buffer.WriteString(pen.String())
				buffer.WriteRune(' ')
			}
			pen.currentVector.X += pen.direction.X * float64(inst.val)
			pen.currentVector.Y += pen.direction.Y * float64(inst.val)
			if pen.down {
				buffer.WriteString(pen.String())
				buffer.WriteRune('\n')
			}
		case "BACK":
			if pen.down {
				buffer.WriteString(pen.color)
				buffer.WriteRune(' ')
				buffer.WriteString(pen.String())
				buffer.WriteRune(' ')
			}
			pen.currentVector.X -= pen.direction.X * float64(inst.val)
			pen.currentVector.Y -= pen.direction.Y * float64(inst.val)
			if pen.down {
				buffer.WriteString(pen.String())
				buffer.WriteRune('\n')
			}
		case "UP":
			pen.down = false
		case "DOWN":
			pen.down = true
		case "LEFT":
			deg := inst.val
			rad := float64(deg) * math.Pi / 180.0
			pen.direction.rotateLeft(rad)
		case "RIGHT":
			deg := inst.val
			rad := float64(deg) * math.Pi / 180.0
			pen.direction.rotateRight(rad)
		case "COLOR":
			pen.color = inst.color
		case "REP":
			n := inst.val
			for ; n > 0; n-- {
				evaluateProgram(inst.instructions, buffer, pen)
			}
		}
	}
	return buffer.String()
}

func syntaxError(line int) error {
	return fmt.Errorf("Syntaxfel på rad %d", line)
}
