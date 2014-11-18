package main

import (
	"github.com/christopherL91/Parser/toki"
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"testing"
)

func TestExample1(t *testing.T) {
	expected := []toki.Token{
		COMMENT,
		COMMENT,
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
	s.SetInput(string(input))
	Convey("Testing that example1.txt is a valid program", t, func() {
		for _, e := range expected {
			r := s.Next()
			if e != r.Token {
				t.Fatalf("expected %v, got %v", e, r.Token)
			} else {
				t.Log(r)
			}
		}
	})
}

func TestExample2(t *testing.T) {
	expected := []toki.Token{
		COMMENT,
		DOWN,
		DOT,
		UP,
		DOT,
		DOWN,
		DOT,
		DOWN,
		DOT,
		COMMENT,
		COMMENT,
		COMMENT,
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
		COMMENT,
		COLORKEYWORD,
		COMMENT,
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
	s.SetInput(string(input))
	Convey("Testing that example2.txt is a valid program", t, func() {
		for _, e := range expected {
			r := s.Next()
			if e != r.Token {
				t.Fatalf("expected %v, got %v", e, r.Token)
			} else {
				t.Log(r)
			}
		}
	})
}
