package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func (g *Game) spawnGhosts(win *pixelgl.Window, sprites []*pixel.Sprite, amount int) []*Ghost {
	ghosts := make([]*Ghost, amount)
	for i := 0; i < amount; i++ {
		x := float64(RandIntInRange(int(win.Bounds().Min.X+16), int(win.Bounds().Max.X-16)))
		y := float64(RandIntInRange(int(win.Bounds().Min.Y+60), int(win.Bounds().Max.Y-16)))
		ttlSec := 5
		ghosts[i] = NewGhost(g, x, y, sprites, ttlSec)
	}
	return ghosts
}

func (g *Game) spawnPacs(win *pixelgl.Window, sprites []*pixel.Sprite, amount int) []*Pac {
	pacs := make([]*Pac, amount)
	for i := 0; i < amount; i++ {
		x := float64(RandIntInRange(int(win.Bounds().Min.X+WindowBoundaryDelta), int(win.Bounds().Max.X-WindowBoundaryDelta)))
		y := float64(RandIntInRange(int(win.Bounds().Min.Y+PanelBoundaryY), int(win.Bounds().Max.Y-WindowBoundaryDelta)))
		ttlSec := 30
		pacs[i] = NewPac(x, y, sprites, ttlSec)
	}
	return pacs
}

func (g *Game) spawnShip(win *pixelgl.Window, sprites []*pixel.Sprite) *Ship {
	ship := NewShip(win.Bounds().Center().X, win.Bounds().Center().Y, sprites)
	return ship
}
