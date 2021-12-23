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

func spawnGhosts(win *pixelgl.Window, sprites []*pixel.Sprite, amount int) []*Ghost {
	ghosts := make([]*Ghost, amount)
	for i := 0; i < amount; i++ {
		x := float64(RandIntInRange(int(win.Bounds().Min.X), int(win.Bounds().Max.X)))
		y := float64(RandIntInRange(int(win.Bounds().Min.Y), int(win.Bounds().Max.Y)))
		ghosts[i] = NewGhost(x, y, sprites)
	}
	return ghosts
}

type Game struct {
	win          *pixelgl.Window
	shotSprites  []*pixel.Sprite
	ghostSprites []*pixel.Sprite
	shots        []*Shot
	ghosts       []*Ghost
}

func NewGame(win *pixelgl.Window) *Game {
	game := &Game{
		win: win,
	}

	game.loadSprites()
	game.ghosts = spawnGhosts(win, game.ghostSprites, 10)

	return game
}

func (f *Game) Update() {
	i := len(f.ghosts) - 1
	for i >= 0 {
		f.ghosts[i].Update()
		if f.ghosts[i].IsDead() {
			f.ghosts = append(f.ghosts[:i], f.ghosts[i+1:]...)
		}
		i--
	}

	i = len(f.shots) - 1
	for i >= 0 {
		f.shots[i].Update()
		if f.shots[i].IsDead() {
			f.shots = append(f.shots[:i], f.shots[i+1:]...)
		}
		i--
	}
}

func (f *Game) AddShot(pos pixel.Vec) {
	shot := NewShot(pos.X, pos.Y, f.shotSprites)
	f.shots = append(f.shots, shot)
}

func (f *Game) Draw(win *pixelgl.Window) {
	for _, ghost := range f.ghosts {
		ghost.Draw(win)
	}
	for _, shot := range f.shots {
		shot.Draw(win)
	}
}
