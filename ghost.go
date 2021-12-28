package main

import (
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Ghost struct {
	_createdAt time.Time
	_dead      bool
	game       *Game
	x          float64
	y          float64
	direction  pixel.Vec
	sprites    []*pixel.Sprite
	matrix     pixel.Matrix
	animation  *Animation
	ttlSec     int
}

func NewGhost(game *Game, x, y float64, sprites []*pixel.Sprite, ttlSec int) *Ghost {
	matrix := pixel.IM.Scaled(pixel.ZV, 1.5).Moved(pixel.V(x, y))
	animation := NewAnimation(100*time.Millisecond, sprites, true)
	now := time.Now()

	speed := -0.01 // RandFloatInRange(-0.002, -0.001)
	direction := pixel.V(x, y).Sub(pixel.V(game.win.Bounds().Center().X, game.win.Bounds().Center().Y)).Scaled(speed)

	// XXX: For now TTL is high, seeing that they point to the ship
	return &Ghost{now, false, game, x, y, direction, sprites, matrix, animation, ttlSec + 999}
}

func (g *Ghost) Draw(win *pixelgl.Window) {
	sprite := g.animation.GetCurrentSprite()
	sprite.Draw(win, g.matrix)
}

func (g *Ghost) Update() {
	g.animation.Update()

	newVec := pixel.V(g.x, g.y).Add(g.direction)
	g.x = newVec.X
	g.y = newVec.Y
	g.matrix = g.matrix.Moved(pixel.V(g.direction.X, g.direction.Y))

	if g._createdAt.Add(time.Duration(g.ttlSec) * time.Second).Before(time.Now()) {
		g._dead = true
	}
}

func (g *Ghost) Kill() {
	g._dead = true
}

func (g *Ghost) IsDead() bool {
	return g._dead
}
