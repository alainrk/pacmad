package main

import (
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

	frames := []pixel.Rect{}
	shotSprites := []*pixel.Sprite{}
	for x := shotSpritesheet.Bounds().Min.X; x < shotSpritesheet.Bounds().Max.X; x += 32 {
		for y := shotSpritesheet.Bounds().Min.Y; y < shotSpritesheet.Bounds().Max.Y; y += 32 {
			frames = append(frames, pixel.R(x, y, x+32, y+32))
			shotSprites = append(shotSprites, pixel.NewSprite(shotSpritesheet, frames[len(frames)-1]))
		}
	}

	frames = []pixel.Rect{}
	ghostSprites := []*pixel.Sprite{}

	startX := ghostSpritesheet.Bounds().Max.X / 3 * 2
	startY := ghostSpritesheet.Bounds().Max.Y / 16 * 4
	endX := ghostSpritesheet.Bounds().Max.X
	endY := ghostSpritesheet.Bounds().Max.Y
	step := 16.0

	// for x := ghostSpritesheet.Bounds().Min.X; x < ghostSpritesheet.Bounds().Max.X; x += 32 {
	// 	for y := ghostSpritesheet.Bounds().Min.Y; y < ghostSpritesheet.Bounds().Max.Y; y += 32 {
	for x := startX; x < endX; x += step {
		for y := startY; y < endY; y += step {
			frames = append(frames, pixel.R(x, y, x+step, y+step))
			ghostSprites = append(ghostSprites, pixel.NewSprite(ghostSpritesheet, frames[len(frames)-1]))
		}
	}

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
		mat = mat.Moved(pixel.V(float64(i*32), float64(i*32)))
		ghost.Draw(win, mat)
	}
}

func (f *Game) Draw(win *pixelgl.Window) {
	f.drawPacmanSpritesForTestJustRemoveThisFunction(win)

	for _, shot := range f.shots {
		shot.Draw(win)
	}
}
