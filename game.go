package main

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

type Game struct {
	points       int
	level        int
	lives        int
	status       string
	paused       bool
	win          *pixelgl.Window
	ghostSprites map[string][]*pixel.Sprite
	shotSprites  []*pixel.Sprite
	pacSprites   []*pixel.Sprite
	shipSprites  []*pixel.Sprite
	bgSprites    []*pixel.Sprite
	shots        []*Shot
	ghosts       []*Ghost
	pacs         []*Pac
	ship         *Ship
	panel        *Panel
}

func (g *Game) init() {
	g.points = 0
	g.level = 1
	g.lives = startLives
	g.paused = false
	g.status = "play"

	g.loadSprites()
	g.ship = g.spawnShip(g.win, g.shipSprites)
	go g.spawnGhostsRoutine()
	go g.spawnPacsRoutine()

	g.panel = NewPanel(PanelBoundaryY, g, g.win)
}

func NewGame(win *pixelgl.Window) *Game {
	g := &Game{
		win: win,
	}
	g.init()
	return g
}

func (f *Game) AddShot(pos pixel.Vec, dir pixel.Vec) {
	shot := NewShot(pos.X, pos.Y, dir.X, dir.Y, f.shotSprites)
	f.shots = append(f.shots, shot)
}

func (g *Game) Update() {
	if g.lives <= 0 {
		g.status = "gameover"
		return
	}

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

func (g *Game) drawGameOver() {
	g.win.Clear(colornames.Red)
	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	basicTxt := text.New(pixel.V(g.win.Bounds().Center().X-120, g.win.Bounds().Center().Y+30), basicAtlas)

	fmt.Fprintf(basicTxt, "GAME OVER\n")
	fmt.Fprintf(basicTxt, "SCORE: %d\n\n", g.points)
	fmt.Fprintf(basicTxt, "Insert Coin")
	basicTxt.Draw(g.win, pixel.IM.Scaled(basicTxt.Orig, 3))
}

func (g *Game) Draw() {
	if g.status == "gameover" {
		g.drawGameOver()
		return
	}

	g.bgSprites[0].Draw(g.win, pixel.IM.Moved(pixel.V(g.win.Bounds().Center().X, g.win.Bounds().Center().Y)))

	for _, ghost := range g.ghosts {
		ghost.Draw(g.win)
	}
	for _, shot := range g.shots {
		shot.Draw(g.win)
	}
	for _, pac := range g.pacs {
		pac.Draw(g.win)
	}

	g.ship.Draw(g.win)
	g.panel.Draw()
}
