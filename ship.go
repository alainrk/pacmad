package main

import (
	"math"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Ship struct {
	x         float64
	y         float64
	direction pixel.Vec
	sprites   []*pixel.Sprite
	matrix    pixel.Matrix
	animation *Animation
}

func NewShip(x, y float64, sprites []*pixel.Sprite) *Ship {
	animation := NewAnimation(100*time.Millisecond, sprites, true)
	direction := pixel.V(0, 0)
	matrix := pixel.IM.Scaled(pixel.ZV, 1).Moved(pixel.V(x, y))
	return &Ship{x, y, direction, sprites, matrix, animation}
}

func (g *Ship) Draw(win *pixelgl.Window) {
	sprite := g.animation.GetCurrentSprite()
	sprite.Draw(win, g.matrix)
}

func (g *Ship) Update(direction pixel.Vec) {
	g.animation.Update()

	newv := pixel.V(g.x, g.y).Sub(direction)
	angle := math.Atan2(newv.Y, newv.X) + math.Pi/2 // adjustment for drawing

	g.matrix = pixel.IM.Scaled(pixel.ZV, 1).Rotated(pixel.ZV, angle).Moved(pixel.V(g.x, g.y))
}
