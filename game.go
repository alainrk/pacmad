package main

import (
	"time"

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

func (g *Game) spawnGhostsRoutine() {
	for {
		time.Sleep(1 * time.Second)
		g.ghosts = append(g.ghosts, spawnGhosts(g.win, g.ghostSprites, 1)...)
	}
}

type Game struct {
	win          *pixelgl.Window
	shotSprites  []*pixel.Sprite
	ghostSprites []*pixel.Sprite
	shots        []*Shot
	ghosts       []*Ghost
}

func NewGame(win *pixelgl.Window) *Game {
	g := &Game{
		win: win,
	}

	g.loadSprites()
	go g.spawnGhostsRoutine()

	return g
}

func (g *Game) Update() {
	i := len(g.ghosts) - 1
	for i >= 0 {
		g.ghosts[i].Update()
		if g.ghosts[i].IsDead() {
			g.ghosts = append(g.ghosts[:i], g.ghosts[i+1:]...)
		}
		i--
	}

	i = len(g.shots) - 1
	for i >= 0 {
		g.shots[i].Update()
		if g.shots[i].IsDead() {
			g.shots = append(g.shots[:i], g.shots[i+1:]...)
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
