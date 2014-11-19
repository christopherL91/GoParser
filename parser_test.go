package main

import (
	"github.com/christopherL91/Parser/toki"
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"testing"
)

// Functional program
func TestExample1(t *testing.T) {
	expected := []toki.Token{
		DOWN,
		DOT,
		FORW,
		NUMBER,
		DOT,
		LEFT,
		NUMBER,
		DOT,
		FORW,
		NUMBER,
		DOT,
		LEFT,
		NUMBER,
		DOT,
		FORW,
		NUMBER,
		DOT,
		LEFT,
		NUMBER,
		DOT,
		FORW,
		NUMBER,
		DOT,
		LEFT,
		NUMBER,
		DOT,
	}
	input, err := ioutil.ReadFile("example1.txt")
	if err != nil {
		t.Fatal(err)
	}
	s := toki.NewScanner(defintions)
	noComments := removeComments.ReplaceAllString(string(input), "")
	s.SetInput(noComments)
	buffer := []*toki.Result{} // Holds all the tokens.
	Convey("Example1: Tokenize program using lexer", t, func() {
		for r := s.Next(); r.Token != toki.EOF; r = s.Next() {
			buffer = append(buffer, r) // Append new token to list.
		}
		So(buffer, ShouldNotBeEmpty)
		Convey("Should be able to compare result with expected", func() {
			if len(buffer) != len(expected) {
				t.FailNow()
			} else {
				for i := 0; i < len(buffer); i++ {
					if buffer[i].Token != expected[i] {
						t.FailNow()
					}
				}
			}
		})
		Convey("Program should be valid", func() {
			err := validateBuffer(buffer)
			So(err, ShouldBeNil)
		})
	})
}

// Functional program
func TestExample2(t *testing.T) {
	expected := []toki.Token{
		DOWN,
		DOT,
		UP,
		DOT,
		DOWN,
		DOT,
		DOWN,
		DOT,
		REP,
		NUMBER,
		STRING,
		COLORKEYWORD,
		COLOR,
		DOT,
		FORW,
		NUMBER,
		DOT,
		LEFT,
		NUMBER,
		DOT,
		COLORKEYWORD,
		COLOR,
		DOT,
		FORW,
		NUMBER,
		DOT,
		LEFT,
		NUMBER,
		DOT,
		STRING,
		COLORKEYWORD,
		COLOR,
		DOT,
		REP,
		NUMBER,
		BACK,
		NUMBER,
		DOT,
	}
	input, err := ioutil.ReadFile("example2.txt")
	if err != nil {
		t.Fatal(err)
	}
	s := toki.NewScanner(defintions)
	noComments := removeComments.ReplaceAllString(string(input), "")
	s.SetInput(noComments)
	buffer := []*toki.Result{} // Holds all the tokens.
	Convey("Example2: Tokenize program using lexer", t, func() {
		for r := s.Next(); r.Token != toki.EOF; r = s.Next() {
			buffer = append(buffer, r) // Append new token to list.
		}
		So(buffer, ShouldNotBeEmpty)
		Convey("Should be able to compare result with expected", func() {
			if len(buffer) != len(expected) {
				t.FailNow()
			} else {
				for i := 0; i < len(buffer); i++ {
					if buffer[i].Token != expected[i] {
						t.FailNow()
					}
				}
			}
		})
		Convey("Program should be valid", func() {
			err := validateBuffer(buffer)
			So(err, ShouldBeNil)
		})
	})
}

// Broken program
func TestExample3(t *testing.T) {
	expected := []toki.Token{
		COLORKEYWORD,
		NUMBER,
		FORW,
		NUMBER,
		DOT,
	}
	input, err := ioutil.ReadFile("example3.txt")
	if err != nil {
		t.Fatal(err)
	}
	s := toki.NewScanner(defintions)
	noComments := removeComments.ReplaceAllString(string(input), "")
	s.SetInput(noComments)
	buffer := []*toki.Result{} // Holds all the tokens.
	Convey("Example3: Tokenize program using lexer", t, func() {
		for r := s.Next(); r.Token != toki.EOF; r = s.Next() {
			buffer = append(buffer, r) // Append new token to list.
		}
		t.Log("Buffer:", buffer)
		So(buffer, ShouldNotBeEmpty)
		t.Log("Length of buffer", len(buffer))
		Convey("Should be able to compare result with expected", func() {
			if len(buffer) != len(expected) {
				t.FailNow()
			} else {
				for i := 0; i < len(buffer); i++ {
					if buffer[i].Token != expected[i] {
						t.FailNow()
					}
				}
			}
		})
		Convey("Program should be valid", func() {
			err := validateBuffer(buffer)
			So(err, ShouldNotBeNil)
		})
	})
}
