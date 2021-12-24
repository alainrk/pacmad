package main

import (
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Shot struct {
	direction pixel.Vec
	x         float64
	y         float64
	destX     float64
	destY     float64
	sprites   []*pixel.Sprite
	matrix    pixel.Matrix
	animation *Animation
	_dead     bool
}

func NewShot(x, y, destX, destY float64, sprites []*pixel.Sprite) *Shot {
	matrix := pixel.IM.Scaled(pixel.ZV, 0.6).Moved(pixel.V(x, y))
	animation := NewAnimation(25*time.Millisecond, sprites, true)

	velocity := -0.05
	direction := pixel.V(x, y).Sub(pixel.V(destX, destY)).Scaled(velocity)

	return &Shot{direction, x, y, destX, destY, sprites, matrix, animation, false}
}

func (s *Shot) Draw(win *pixelgl.Window) {
	sprite := s.animation.GetCurrentSprite()
	sprite.Draw(win, s.matrix)
}

func (s *Shot) Update(win *pixelgl.Window) {
	s.animation.Update()

	newVec := pixel.V(s.x, s.y).Add(s.direction)
	s.x = newVec.X
	s.y = newVec.Y
	s.matrix = s.matrix.Moved(s.direction)

	if s.x < win.Bounds().Min.X || s.x > win.Bounds().Max.X || s.y < win.Bounds().Min.Y || s.y > win.Bounds().Max.Y {
		s._dead = true
	}
}

func (s *Shot) Destroy() {
	s._dead = true
}

func (s *Shot) IsDead() bool {
	return s._dead
}
