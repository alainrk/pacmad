package main

import (
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const ghostStageDuration = 35 * time.Millisecond

type Ghost struct {
	x         float64
	y         float64
	createdAt time.Time
	lastStage time.Time
	sprites   []*pixel.Sprite
	matrix    pixel.Matrix
	stage     int
	dead      bool
}

func NewGhost(x float64, y float64, sprites []*pixel.Sprite) *Ghost {
	matrix := pixel.IM.Scaled(pixel.ZV, 1.5).Moved(pixel.Vec{x, y})
	now := time.Now()
	ghost := &Ghost{x, y, now, now, sprites, matrix, 0, false}
	return ghost
}

func (s *Ghost) Draw(win *pixelgl.Window) {
	sprite := s.sprites[s.stage]
	sprite.Draw(win, s.matrix)
}

func (s *Ghost) Update() {
	if s.dead {
		return
	}

	now := time.Now()
	if now.Sub(s.lastStage) >= ghostStageDuration {
		s.stage++
		s.lastStage = time.Now()
	}

	if s.stage >= len(s.sprites) {
		s.dead = true
	}
}
