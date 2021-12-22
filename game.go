package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Game struct {
	shotSprites []*pixel.Sprite
	// ghostSprites []*pixel.Sprite
	shots []*Shot
}

func NewGame() *Game {
	spritesheet, err := loadPicture("shot.png")
	if err != nil {
		panic(err)
	}

	frames := []pixel.Rect{}
	shotSprites := []*pixel.Sprite{}
	for x := spritesheet.Bounds().Min.X; x < spritesheet.Bounds().Max.X; x += 32 {
		for y := spritesheet.Bounds().Min.Y; y < spritesheet.Bounds().Max.Y; y += 32 {
			frames = append(frames, pixel.R(x, y, x+32, y+32))
			shotSprites = append(shotSprites, pixel.NewSprite(spritesheet, frames[len(frames)-1]))
		}
	}

	return &Game{
		shotSprites: shotSprites,
		shots:       []*Shot{},
	}
}

func (f *Game) Update() {
	for i, shot := range f.shots {
		shot.Update()
		if shot.dead {
			f.shots = append(f.shots[:i], f.shots[i+1:]...)
		}
	}
}

func (f *Game) AddShot(pos pixel.Vec) {
	shot := NewShot(pos.X, pos.Y, f.shotSprites)

	f.shots = append(f.shots, shot)
}

func (f *Game) Draw(win *pixelgl.Window) {
	for _, shot := range f.shots {
		shot.Draw(win)
	}
}
