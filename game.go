package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Game struct {
	points       int
	level        int
	lives        int
	status       string
	paused       bool
	win          *pixelgl.Window
	shotSprites  []*pixel.Sprite
	ghostSprites []*pixel.Sprite
	pacSprites   []*pixel.Sprite
	shipSprites  []*pixel.Sprite
	shots        []*Shot
	ghosts       []*Ghost
	pacs         []*Pac
	ship         *Ship
	panel        *Panel
}

func NewGame(win *pixelgl.Window) *Game {
	g := &Game{
		points: 0,
		level:  1,
		lives:  3,
		paused: false,
		status: "play",
		win:    win,
	}

	g.loadSprites()
	g.ship = g.spawnShip(win, g.shipSprites)
	go g.spawnGhostsRoutine()
	go g.spawnPacsRoutine()

	g.panel = NewPanel(PanelBoundaryY, g, win)
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
	g.panel.Draw()
}
