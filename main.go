// +build js,wasm

package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fractalbach/go-wasm-exp/pkg/canvas"
	"github.com/fractalbach/go-wasm-exp/pkg/shape"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// =======================================================================

type Game struct {
	objects []*GameObject
	w       float64
	h       float64
}

func (g *Game) step() {
	for _, o := range g.objects {
		nextX := o.X + o.Vx
		nextY := o.Y + o.Vy
		nextVx := o.Vx
		nextVy := o.Vy
		if nextX < 0 {
			nextX = -nextX
			nextVx *= -1
		}
		if nextX > g.w {
			nextX = g.w - 2*(nextX-g.w)
			nextVx *= -1
		}
		if nextY < 0 {
			nextY = -nextY
			nextVy *= -1
		}
		if nextY > g.h {
			nextY = g.h - 2*(nextY-g.h)
			nextVy *= -1
		}
		o.X = nextX
		o.Y = nextY
		o.Vx = nextVx
		o.Vy = nextVy
	}
}

// =======================================================================

type GameObject struct {
	shape.Circle
	Vx float64
	Vy float64
}

// =======================================================================

func randFloat(min, max float64) float64 {
	return min + r.Float64()*(max-min)
}

func randGameObject(maxX, maxY float64) *GameObject {
	r := randFloat(10, 20)
	return &GameObject{
		Circle: shape.Circle{
			X: randFloat(r/2, (maxX)-r/2),
			Y: randFloat(r/2, (maxY)-r/2),
			R: r,
		},
		Vx: randFloat(-4, 4),
		Vy: randFloat(-4, 4),
	}
}

func print(args ...interface{}) {
	fmt.Println("WasmBach:" + fmt.Sprint(args...))
}

// =======================================================================

func main() {

	print("main started.")
	defer print("main ended.")

	const (
		n = 30
		w = 1000.0
		h = 1000.0
	)

	g := &Game{
		objects: make([]*GameObject, n),
		w:       w,
		h:       h,
	}

	// populate game with objects.
	for i := 0; i < n; i++ {
		g.objects[i] = randGameObject(g.w, g.h)
	}

	runForWeb(g)
}

// ========================================================================
// build for web only
// TODO: move to separate package

func runForWeb(game *Game) {
	var (
		ping    = make(chan bool)
		pong    = make(chan bool)
		counter = 0
		c       = canvas.New(game.w, game.h)
	)
	draw := func() {
		for _, obj := range game.objects {
			c.DrawCircle(&obj.Circle)
		}
	}
	canvas.AnimationLoop(ping, pong, draw)
	for {
		select {
		case <-ping:
			c.Clear()
			game.step()
			draw()
			counter++
			if counter%1000 == 0 {
				fmt.Println("counter=", counter, ", time=", time.Now())
			}
			pong <- true
		}
	}
}
