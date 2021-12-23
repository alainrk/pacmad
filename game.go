package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func createSprites(spritesheet *pixel.Picture, minX, minY, maxX, maxY, step float64) []*pixel.Sprite {
	sprites := []*pixel.Sprite{}

	for y := minY; y < maxY; y += step {
		for x := minX; x < maxX; x += step {
			frame := pixel.R(x, y, x+step, y+step)
			sprites = append(sprites, pixel.NewSprite(*spritesheet, frame))
		}
	}

	return sprites
}

func (f *Game) loadShotSprites() {
	spritesheet, err := loadPicture("shot.png")
	if err != nil {
		panic(err)
	}

	startX := spritesheet.Bounds().Min.X
	startY := spritesheet.Bounds().Min.Y
	endX := spritesheet.Bounds().Max.X
	endY := spritesheet.Bounds().Max.Y

	f.shotSprites = createSprites(&spritesheet, startX, startY, endX, endY, 32)
}

func (f *Game) loadGhostSprites() {
	spritesheet, err := loadPicture("pmsprites.png")
	if err != nil {
		panic(err)
	}

	step := 16.0
	startX := spritesheet.Bounds().Max.X/3*2 + 3
	startY := 168.0
	endX := startX + (step * 8)
	endY := startY + step

	f.ghostSprites = createSprites(&spritesheet, startX, startY, endX, endY, step)
}

func (f *Game) loadSprites() {
	f.loadGhostSprites()
	f.loadShotSprites()
}

type Game struct {
	shotSprites  []*pixel.Sprite
	ghostSprites []*pixel.Sprite
	shots        []*Shot
	ghosts       []*Ghost
}

func NewGame() *Game {
	game := &Game{
		shotSprites:  nil,
		ghostSprites: nil,
		shots:        []*Shot{},
		ghosts:       []*Ghost{},
	}

	game.loadSprites()

	return game
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
	f.drawPacmanSpritesForTestJustRemoveThisFunction(win)

	for _, shot := range f.shots {
		shot.Draw(win)
	}
}
