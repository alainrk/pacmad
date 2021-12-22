package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Pac struct {
	x      float64
	y      float64
	angle  float64
	sprite *pixel.Sprite
}

func (p *Pac) move(dx, dy, maxx, maxy float64) {
	p.x += dx
	p.y += dy
}

func NewPac() *Pac {
	pic, err := loadPicture("pac.png")
	if err != nil {
		panic(err)
	}

	sprite := pixel.NewSprite(pic, pic.Bounds())

	return &Pac{
		x:      0,
		y:      0,
		angle:  0,
		sprite: sprite,
	}
}

func (p *Pac) Move(dt float64) {
	p.angle += 5 * dt
	p.x += 1000 * dt
	p.y += 1000 * dt
}

func (p *Pac) Draw(win *pixelgl.Window) {
	mat := pixel.IM
	mat = mat.Moved(pixel.V(p.x, p.y))
	mat = mat.ScaledXY(win.Bounds().Center(), pixel.V(0.02, 0.02))
	p.sprite.Draw(win, mat)
}
