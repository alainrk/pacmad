package main

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
)

type Game struct {
	points       int
	win          *pixelgl.Window
	shotSprites  []*pixel.Sprite
	ghostSprites []*pixel.Sprite
	pacSprites   []*pixel.Sprite
	shipSprites  []*pixel.Sprite
	shots        []*Shot
	ghosts       []*Ghost
	pacs         []*Pac
	ship         *Ship
}

func NewGame(win *pixelgl.Window) *Game {
	g := &Game{
		points: 0,
		win:    win,
	}

	g.loadSprites()
	g.ship = spawnShip(win, g.shipSprites)
	go g.spawnGhostsRoutine()
	go g.spawnPacsRoutine()

	return g
}

func (f *Game) AddShot(pos pixel.Vec, dir pixel.Vec) {
	shot := NewShot(pos.X, pos.Y, dir.X, dir.Y, f.shotSprites)
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
		g.shots[i].Update(g.win)
		if g.shots[i].IsDead() {
			g.shots = append(g.shots[:i], g.shots[i+1:]...)
		}
		i--
	}

	g.ship.Update(g.win.MousePosition())
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

	g.ship.Draw(win)

	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	basicTxt := text.New(pixel.V(20, 30), basicAtlas)

	fmt.Fprintf(basicTxt, "SCORE: %d\n", g.points)
	fmt.Fprintln(basicTxt, "Click to shoot")
	basicTxt.Draw(win, pixel.IM.Scaled(basicTxt.Orig, 1.3))
}
