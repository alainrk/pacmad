package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func generateValidSpawningPosition(win *pixelgl.Window) (float64, float64) {
	var x, y float64
	// TODO: This can be evaluated just once
	excludeXRange := [2]float64{win.Bounds().Center().X - SpawnDeltaBoundaryFromCenter, win.Bounds().Center().X + SpawnDeltaBoundaryFromCenter}
	excludeYRange := [2]float64{win.Bounds().Center().Y - SpawnDeltaBoundaryFromCenter, win.Bounds().Center().Y + SpawnDeltaBoundaryFromCenter}
	for {
		x = float64(RandIntInRange(int(win.Bounds().Min.X+WindowBoundaryDelta), int(win.Bounds().Max.X-WindowBoundaryDelta)))
		if x > excludeXRange[0] && x < excludeXRange[1] {
			continue
		}
		break
	}
	for {
		y = float64(RandIntInRange(int(win.Bounds().Min.Y+PanelBoundaryY), int(win.Bounds().Max.Y-WindowBoundaryDelta)))
		if y > excludeYRange[0] && y < excludeYRange[1] {
			continue
		}
		break
	}
	return x, y
}

func (g *Game) spawnGhosts(win *pixelgl.Window, sprites []*pixel.Sprite, amount int) []*Ghost {
	ghosts := make([]*Ghost, amount)
	for i := 0; i < amount; i++ {
		x, y := generateValidSpawningPosition(win)
		ttlSec := 5
		ghosts[i] = NewGhost(g, x, y, sprites, ttlSec)
	}
	return ghosts
}

func (g *Game) spawnPacs(win *pixelgl.Window, sprites []*pixel.Sprite, amount int) []*Pac {
	pacs := make([]*Pac, amount)
	for i := 0; i < amount; i++ {
		x, y := generateValidSpawningPosition(win)
		ttlSec := 30
		pacs[i] = NewPac(x, y, sprites, ttlSec)
	}
	return pacs
}

func (g *Game) spawnShip(win *pixelgl.Window, sprites []*pixel.Sprite) *Ship {
	ship := NewShip(win.Bounds().Center().X, win.Bounds().Center().Y, sprites)
	return ship
}
