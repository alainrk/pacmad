package main

import (
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Ghost struct {
	x         float64
	y         float64
	sprites   []*pixel.Sprite
	matrix    pixel.Matrix
	animation *Animation
}

func NewGhost(x, y float64, sprites []*pixel.Sprite) *Ghost {
	matrix := pixel.IM.Scaled(pixel.ZV, 1.5).Moved(pixel.V(x, y))
	animation := NewAnimation(100*time.Millisecond, sprites, true)
	return &Ghost{x, y, sprites, matrix, animation}
}

func (g *Ghost) Draw(win *pixelgl.Window) {
	sprite := g.animation.GetCurrentSprite()
	sprite.Draw(win, g.matrix)
}

func (g *Ghost) Update() {
	g.animation.Update()
}

func (g *Ghost) IsDead() bool {
	return g.animation.IsDead()
}
