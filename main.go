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

/*
Game is the logical core of the game itself.  All platform-independent game
state and logic are represented.  To port the game to another platform,
Write platform-specific operations that interact with this Game object.
*/
type Game struct {
	objects []*GameObject
	w       float64
	h       float64
}

/*
Step does one single iteration of the game.  New object positions are
calculated in this function.
*/
func (g *Game) Step() {
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

/*
RunGameTicker loops forever, calling game.Step() every timePerTick. Note that
RunGameTicker is a blocking call, so call it using the "go" keyword if you want
to continue execution and let the game ticker run in the background.
*/
func RunGameTicker(game *Game, timePerTick time.Duration) {
	ticker := time.NewTicker(timePerTick)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			game.Step()
		}
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
	// print("main started.")
	// defer js.Global().Call("displayError", "internal error")
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
	runForWeb(g) // TODO: use only in "web-only" mode.
}

// ========================================================================
// build for web only
// TODO: move to separate package

func runForWeb(game *Game) {
	var (
		c           = canvas.New(game.w, game.h)
		timePerTick = time.Millisecond * 20
	)
	draw := func() {
		c.Clear()
		for _, obj := range game.objects {
			c.DrawCircle(&obj.Circle)
		}
	}
	canvas.DefineAnimationLoop(draw)
	RunGameTicker(game, timePerTick)
}
