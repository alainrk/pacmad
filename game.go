package main

import (
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const shotDuration = 100 * time.Millisecond

type Game struct {
	spritesheet     pixel.Picture
	availableFrames []pixel.Rect
	sprites         []*pixel.Sprite
	shots           []*Shot
}

func NewGame() *Game {
	spritesheet, err := loadPicture("shot.png")
	if err != nil {
		panic(err)
	}

	frames := []pixel.Rect{}
	sprites := []*pixel.Sprite{}
	for x := spritesheet.Bounds().Min.X; x < spritesheet.Bounds().Max.X; x += 32 {
		for y := spritesheet.Bounds().Min.Y; y < spritesheet.Bounds().Max.Y; y += 32 {
			frames = append(frames, pixel.R(x, y, x+32, y+32))
			sprites = append(sprites, pixel.NewSprite(spritesheet, frames[len(frames)-1]))
		}
	}

	return &Game{
		spritesheet:     spritesheet,
		availableFrames: frames,
		sprites:         sprites,
		shots:           []*Shot{},
	}
}

func (f *Game) Update() {
	now := time.Now()
	for i, shot := range f.shots {
		if now.Sub(shot.createdAt) > shotDuration {
			f.shots = append(f.shots[:i], f.shots[i+1:]...)
		}
	}
}

func (f *Game) AddShot(pos pixel.Vec) {
	shot := NewShot(pos.X, pos.Y, f.sprites)

	f.shots = append(f.shots, shot)
}

func (f *Game) Draw(win *pixelgl.Window) {
	for _, shot := range f.shots {
		shot.Draw(win)
	}
}
