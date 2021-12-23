package main

import (
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Shot struct {
	x         float64
	y         float64
	sprites   []*pixel.Sprite
	matrix    pixel.Matrix
	animation *Animation
}

func NewShot(x, y float64, sprites []*pixel.Sprite) *Shot {
	matrix := pixel.IM.Scaled(pixel.ZV, 1.2).Moved(pixel.V(x, y))
	animation := NewAnimation(35*time.Millisecond, sprites, false)
	return &Shot{x, y, sprites, matrix, animation}
}

func (s *Shot) Draw(win *pixelgl.Window) {
	sprite := s.animation.GetCurrentSprite()
	sprite.Draw(win, s.matrix)
}

func (s *Shot) Update() {
	s.animation.Update()
}

func (s *Shot) IsDead() bool {
	return s.animation.IsDead()
}
