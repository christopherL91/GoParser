/*
	Author: Christopher Lillthors
	License: MIT
*/

package main

import (
	"fmt"
	"math"
)

type Vector struct {
	X, Y float64
}

type Pen struct {
	down          bool
	color         string
	currentVector *Vector
	direction     *Vector
}

func newPen() *Pen {
	return &Pen{
		down:          false,
		color:         "#0000FF",
		currentVector: newVektor(0, 0),
		direction:     newVektor(1, 0),
	}
}

func (p *Pen) String() string {
	return fmt.Sprintf("%.4f %.4f", p.currentVector.X, p.currentVector.Y)
}

/*
	Creates a new vector.
*/
func newVektor(x, y float64) *Vector {
	return &Vector{x, y}
}

/*
	Rotates a vector by rot degrees.
*/
func (p *Vector) rotateLeft(rad float64) {
	x := p.X
	y := p.Y
	p.X = x*math.Cos(rad) - y*math.Sin(rad)
	p.Y = x*math.Sin(rad) + y*math.Cos(rad)
}

func (p *Vector) rotateRight(rad float64) {
	x := p.X
	y := p.Y
	p.X = x*math.Cos(rad) + y*math.Sin(rad)
	p.Y = -x*math.Sin(rad) + y*math.Cos(rad)
}
