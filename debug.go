package main

import (
	"fmt"
	"strings"

	"github.com/christopherL91/Parser/toki"
)

/*
	Used for debugging purposes.
*/
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
