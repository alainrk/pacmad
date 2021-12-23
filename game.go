package main

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Game struct {
	shotSprites  []*pixel.Sprite
	ghostSprites []*pixel.Sprite
	shots        []*Shot
}

func NewGame() *Game {
	shotSpritesheet, err := loadPicture("shot.png")
	if err != nil {
		panic(err)
	}
	ghostSpritesheet, err := loadPicture("pmsprites.png")
	if err != nil {
		panic(err)
	}

	shotSprites := []*pixel.Sprite{}
	for y := shotSpritesheet.Bounds().Min.Y; y < shotSpritesheet.Bounds().Max.Y; y += 32 {
		for x := shotSpritesheet.Bounds().Min.X; x < shotSpritesheet.Bounds().Max.X; x += 32 {
			frame := pixel.R(x, y, x+32, y+32)
			shotSprites = append(shotSprites, pixel.NewSprite(shotSpritesheet, frame))
		}
	}

	ghostSprites := []*pixel.Sprite{}

	step := 16.0
	startX := ghostSpritesheet.Bounds().Max.X/3*2 + 3
	startY := 168.0
	endX := startX + (step * 8)
	endY := startY + step

	for y := startY; y < endY; y += step {
		for x := startX; x < endX; x += step {
			frame := pixel.R(x, y, x+step, y+step)
			ghostSprites = append(ghostSprites, pixel.NewSprite(ghostSpritesheet, frame))
		}
	}
	fmt.Println(len(ghostSprites))

	return &Game{
		shotSprites:  shotSprites,
		ghostSprites: ghostSprites,
		shots:        []*Shot{},
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

func (f *Game) drawPacmanSpritesForTestJustRemoveThisFunction(win *pixelgl.Window) {
	for i, ghost := range f.ghostSprites {
		mat := pixel.IM
		mat = mat.Moved(pixel.V(float64(i*16)+float64(i*2.0), 0))
		ghost.Draw(win, mat)
	}
}

func (f *Game) Draw(win *pixelgl.Window) {
	// f.drawPacmanSpritesForTestJustRemoveThisFunction(win)

	for _, shot := range f.shots {
		shot.Draw(win)
	}
}
