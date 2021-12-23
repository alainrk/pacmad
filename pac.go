package main

import (
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Pac struct {
	_createdAt time.Time
	_dead      bool
	x          float64
	y          float64
	direction  pixel.Vec
	sprites    []*pixel.Sprite
	matrix     pixel.Matrix
	animation  *Animation
	ttlSec     int
}

func NewPac(x, y float64, sprites []*pixel.Sprite, ttlSec int) *Pac {
	matrix := pixel.IM.Scaled(pixel.ZV, 1).Moved(pixel.V(x, y))
	animation := NewAnimation(100*time.Millisecond, sprites, true)
	now := time.Now()
	dx, dy := RandFloatInRange(-1, 1), RandFloatInRange(-1, 1)
	direction := pixel.V(dx, dy)
	return &Pac{now, false, x, y, direction, sprites, matrix, animation, ttlSec}
}

func (p *Pac) Draw(win *pixelgl.Window) {
	sprite := p.animation.GetCurrentSprite()
	sprite.Draw(win, p.matrix)
}

func (p *Pac) Update() {
	p.animation.Update()

	newVec := pixel.V(p.x, p.y).Add(p.direction)
	p.x = newVec.X
	p.y = newVec.Y
	p.matrix = p.matrix.Moved(pixel.V(p.direction.X, p.direction.Y))

	if p._createdAt.Add(time.Duration(p.ttlSec) * time.Second).Before(time.Now()) {
		p._dead = true
	}
}

func (p *Pac) Kill() {
	p._dead = true
}

func (p *Pac) IsDead() bool {
	return p._dead
}
