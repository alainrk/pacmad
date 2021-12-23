package main

import (
	"time"

	"github.com/faiface/pixel"
)

type Animation struct {
	_createdAt    time.Time
	_lastStage    time.Time
	_stage        int
	_dead         bool
	stageDuration time.Duration
	sprites       []*pixel.Sprite
	loop          bool
}

func NewAnimation(stageDuration time.Duration, sprites []*pixel.Sprite, loop bool) *Animation {
	animation := &Animation{
		_createdAt:    time.Now(),
		_lastStage:    time.Now(),
		_stage:        0,
		stageDuration: stageDuration,
		sprites:       sprites,
		loop:          loop,
	}
	return animation
}

func (a *Animation) CurrentSprite() *pixel.Sprite {
	return a.sprites[a._stage]
}

func (a *Animation) IsDead() bool {
	return a._dead
}

func (a *Animation) Update() {
	if a._dead {
		return
	}

	now := time.Now()
	if now.Sub(a._lastStage) >= a.stageDuration {
		a._stage++
		a._lastStage = time.Now()
	}

	if a._stage >= len(a.sprites) {
		a._dead = true
	}
}
