package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

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

func spawnShip(win *pixelgl.Window, sprites []*pixel.Sprite) *Ship {
	ship := NewShip(win.Bounds().Center().X, win.Bounds().Center().Y, sprites)
	return ship
}