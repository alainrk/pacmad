package main

import (
	"fmt"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
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

func spawnGhosts(win *pixelgl.Window, sprites []*pixel.Sprite, amount int) []*Ghost {
	ghosts := make([]*Ghost, amount)
	for i := 0; i < amount; i++ {
		x := float64(RandIntInRange(int(win.Bounds().Min.X+16), int(win.Bounds().Max.X-16)))
		y := float64(RandIntInRange(int(win.Bounds().Min.Y+60), int(win.Bounds().Max.Y-16)))
		ttlSec := 5
		ghosts[i] = NewGhost(x, y, sprites, ttlSec)
	}
	return ghosts
}

func spawnPacs(win *pixelgl.Window, sprites []*pixel.Sprite, amount int) []*Pac {
	pacs := make([]*Pac, amount)
	for i := 0; i < amount; i++ {
		x := float64(RandIntInRange(int(win.Bounds().Min.X+WindowBoundaryDelta), int(win.Bounds().Max.X-WindowBoundaryDelta)))
		y := float64(RandIntInRange(int(win.Bounds().Min.Y+WindowBoundaryDeltaY), int(win.Bounds().Max.Y-WindowBoundaryDelta)))
		ttlSec := 30
		pacs[i] = NewPac(x, y, sprites, ttlSec)
	}
	return pacs
}

type Game struct {
	points       int
	win          *pixelgl.Window
	shotSprites  []*pixel.Sprite
	ghostSprites []*pixel.Sprite
	pacSprites   []*pixel.Sprite
	shots        []*Shot
	ghosts       []*Ghost
	pacs         []*Pac
}

func NewGame(win *pixelgl.Window) *Game {
	g := &Game{
		points: 0,
		win:    win,
	}

	g.loadSprites()
	go g.spawnGhostsRoutine()
	go g.spawnPacsRoutine()

	return g
}

func checkCollision(x1, y1, x2, y2 float64, box1w, box1h, box2w, box2h float64) bool {
	return x1 < x2+box2w &&
		x1+box1w > x2 &&
		y1 < y2+box2h &&
		box1h+y1 > y2
}

func (g *Game) resolveCollisions() {
	for _, ghost := range g.ghosts {
		for _, shot := range g.shots {
			if checkCollision(ghost.x, ghost.y, shot.x, shot.y, 16.0, 16.0, 16.0, 16.0) {
				g.points++
				ghost.Kill()
			}
		}
	}

	for _, pac := range g.pacs {
		for _, ghost := range g.ghosts {
			if checkCollision(pac.x, pac.y, ghost.x, ghost.y, 32.0, 32.0, 16.0, 16.0) {
				g.points -= 10
				pac.Kill()
			}
		}
	}

	for _, pac := range g.pacs {
		for _, shot := range g.shots {
			if checkCollision(pac.x, pac.y, shot.x, shot.y, 32.0, 32.0, 16.0, 16.0) {
				g.points += 10
				pac.Kill()
			}
		}
	}
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

func (f *Game) loadPacSprites() {
	spritesheet, err := loadPicture("pac.png")
	if err != nil {
		panic(err)
	}

	startX := spritesheet.Bounds().Min.X
	startY := spritesheet.Bounds().Min.Y
	endX := spritesheet.Bounds().Max.X
	endY := spritesheet.Bounds().Max.Y

	f.pacSprites = createSprites(&spritesheet, startX, startY, endX, endY, 32)
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
	f.loadPacSprites()
}

func (g *Game) spawnGhostsRoutine() {
	for {
		s := RandIntInRange(GhostSpawnIntervalMin, GhostSpawnIntervalMax)
		time.Sleep(time.Duration(s * int(time.Millisecond)))
		g.ghosts = append(g.ghosts, spawnGhosts(g.win, g.ghostSprites, 1)...)
	}
}

func (g *Game) spawnPacsRoutine() {
	for {
		s := RandIntInRange(PacSpawnIntervalMin, PacSpawnIntervalMax)
		time.Sleep(time.Duration(s * int(time.Millisecond)))
		g.pacs = append(g.pacs, spawnPacs(g.win, g.pacSprites, 1)...)
	}
}

func (f *Game) AddShot(pos pixel.Vec) {
	shot := NewShot(pos.X, pos.Y, f.shotSprites)
	f.shots = append(f.shots, shot)
}

func (g *Game) Update() {
	g.resolveCollisions()

	i := len(g.ghosts) - 1
	for i >= 0 {
		g.ghosts[i].Update()
		if g.ghosts[i].IsDead() {
			g.ghosts = append(g.ghosts[:i], g.ghosts[i+1:]...)
		}
		i--
	}

	i = len(g.pacs) - 1
	for i >= 0 {
		g.pacs[i].Update()
		if g.pacs[i].IsDead() {
			g.pacs = append(g.pacs[:i], g.pacs[i+1:]...)
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

func (g *Game) Draw(win *pixelgl.Window) {
	for _, ghost := range g.ghosts {
		ghost.Draw(win)
	}
	for _, shot := range g.shots {
		shot.Draw(win)
	}
	for _, pac := range g.pacs {
		pac.Draw(win)
	}

	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	basicTxt := text.New(pixel.V(20, 30), basicAtlas)

	fmt.Fprintf(basicTxt, "SCORE: %d\n", g.points)
	fmt.Fprintln(basicTxt, "Click to shoot")
	basicTxt.Draw(win, pixel.IM.Scaled(basicTxt.Orig, 1.3))
}
