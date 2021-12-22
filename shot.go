package main

import (
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const shotStageDuration = 35 * time.Millisecond

type Shot struct {
	x         float64
	y         float64
	createdAt time.Time
	lastStage time.Time
	sprites   []*pixel.Sprite
	matrix    pixel.Matrix
	stage     int
	dead      bool
}

func NewShot(x float64, y float64, sprites []*pixel.Sprite) *Shot {
	matrix := pixel.IM.Scaled(pixel.ZV, 1.5).Moved(pixel.Vec{x, y})
	now := time.Now()
	shot := &Shot{x, y, now, now, sprites, matrix, 0, false}
	return shot
}

func (s *Shot) Draw(win *pixelgl.Window) {
	sprite := s.sprites[s.stage]
	sprite.Draw(win, s.matrix)
}

func (s *Shot) Update() {
	if s.dead {
		return
	}

	now := time.Now()
	if now.Sub(s.lastStage) >= shotStageDuration {
		s.stage++
		s.lastStage = time.Now()
	}

	if s.stage >= len(s.sprites) {
		s.dead = true
	}
}
