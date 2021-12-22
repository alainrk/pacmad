package main

import (
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Shot struct {
	x         float64
	y         float64
	createdAt time.Time
	sprites   []*pixel.Sprite
	matrix    pixel.Matrix
	stage     int
}

func (s *Shot) Draw(win *pixelgl.Window) {
	sprite := s.sprites[s.stage]
	sprite.Draw(win, s.matrix)
}

func NewShot(x float64, y float64, sprites []*pixel.Sprite) *Shot {
	matrix := pixel.IM.Scaled(pixel.ZV, 1.5).Moved(pixel.Vec{x, y})
	now := time.Now()
	shot := &Shot{x, y, now, sprites, matrix, 0}

	return shot
}
