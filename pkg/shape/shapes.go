// Package shape implements some primitives for 2d graphics.
package shape

import "fmt"

// Circle is defined by its center (x,y) and radius r.
type Circle struct {
	X float64
	Y float64
	R float64
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle{X:%v, Y:%v, R:%v}", c.X, c.Y, c.R)
}
