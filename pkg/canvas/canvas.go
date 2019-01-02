// +build js,wasm

// Package canvas is an experimental html canvas with 2d drawing abilities.
package canvas

import (
	"math"
	"syscall/js"

	"github.com/fractalbach/go-wasm-exp/pkg/shape"
	// "github.com/dennwc/dom"
)

// Canvas Element represents the HTML5 canvas.
// Specifically designed for 2d graphics.
// This is an experiment with webassembly and drawing graphics.
type Canvas struct {
	width   float64
	height  float64
	element js.Value
	ctx     js.Value
}

// New creates a new Canvas element, and returns its reference.
func New(w, h float64) *Canvas {
	// ele := js.Global().Get("document").Call("createElement", "Canvas")
	ele := js.Global().Get("document").Call("querySelector", "#game")
	ctx := ele.Call("getContext", "2d")
	canv := &Canvas{
		width:   w,
		height:  h,
		element: ele,
		ctx:     ctx,
	}
	// js.Global().Get("document").Get("body").Call("appendChild", canv.element)
	// canv.element.Call("setAttribute", "width", w)
	// canv.element.Call("setAttribute", "height", h)
	canv.element.Get("style").Call("setProperty", "border", "1px solid black")
	return canv
}

func DefineAnimationLoop(f func()) {
	var loop js.Callback
	loop = js.NewCallback(func(args []js.Value) {
		f()
		js.Global().Get("window").Call("requestAnimationFrame", loop)
	})
	js.Global().Get("window").Call("requestAnimationFrame", loop)
	js.Global().Call("doneLoading")
}

// Drawable can be drawn onto a canvas using the Draw function.
// type Drawable interface {
// 	Draw(*Canvas)
// }
//
// // Draw a drawable shape onto the canvas.
// func (c *Canvas) Draw(d Drawable) {
// 	d.Draw(c)
// }

func (canv *Canvas) DrawCircle(c *shape.Circle) {
	canv.ctx.Call("beginPath")
	canv.ctx.Call("arc", c.X, c.Y, c.R, "0", math.Pi*2, "true")
	canv.ctx.Call("stroke")
}

func (canv *Canvas) Clear() {
	canv.ctx.Call("clearRect", 0, 0, canv.width, canv.height)
}
