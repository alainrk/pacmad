package main

import (
	"fmt"

	"github.com/faiface/pixel"
)

func (g *Game) loadBackgroundSprites() {
	spritesheet, err := loadPicture("assets/background.png")
	if err != nil {
		panic(err)
	}

	startX := spritesheet.Bounds().Min.X
	startY := spritesheet.Bounds().Min.Y
	endX := spritesheet.Bounds().Max.X
	endY := spritesheet.Bounds().Max.Y

	g.bgSprites = createSprites(&spritesheet, startX, startY, endX, endY, endX, endY)
	fmt.Println(startX, startY, endX, endY)
}

func (g *Game) loadShotSprites() {
	spritesheet, err := loadPicture("assets/shot.png")
	if err != nil {
		panic(err)
	}

	startX := spritesheet.Bounds().Min.X
	startY := spritesheet.Bounds().Min.Y
	endX := spritesheet.Bounds().Max.X
	endY := spritesheet.Bounds().Max.Y

	g.shotSprites = createSprites(&spritesheet, startX, startY, endX, endY, 32, 32)
}

func (g *Game) loadPacSprites() {
	spritesheet, err := loadPicture("assets/pac.png")
	if err != nil {
		panic(err)
	}

	startX := spritesheet.Bounds().Min.X
	startY := spritesheet.Bounds().Min.Y
	endX := spritesheet.Bounds().Max.X
	endY := spritesheet.Bounds().Max.Y

	g.pacSprites = createSprites(&spritesheet, startX, startY, endX, endY, 32, 32)
}

func (g *Game) loadShipSprites() {
	spritesheet, err := loadPicture("assets/ship.png")
	if err != nil {
		panic(err)
	}

	startX := spritesheet.Bounds().Min.X
	startY := spritesheet.Bounds().Min.Y
	endX := spritesheet.Bounds().Max.X
	endY := spritesheet.Bounds().Max.Y

	g.shipSprites = createSprites(&spritesheet, startX, startY, endX, endY, 32, 32)
}

func (g *Game) loadGhostSprites() {
	spritesheet, err := loadPicture("assets/pmsprites.png")
	if err != nil {
		panic(err)
	}

	ghosts := map[string][]*pixel.Sprite{"orange": nil, "blue": nil, "pink": nil, "red": nil}

	step := 16.0
	startX := spritesheet.Bounds().Max.X/3*2 + 3
	endX := startX + (step * 8)

	startY := float64(120.0)
	endY := startY + step

	ghosts["orange"] = createSprites(&spritesheet, startX, startY+0*step, endX, endY+0*step, 16, 16)
	ghosts["blue"] = createSprites(&spritesheet, startX, startY+1*step, endX, endY+1*step, 16, 16)
	ghosts["pink"] = createSprites(&spritesheet, startX, startY+2*step, endX, endY+2*step, 16, 16)
	ghosts["red"] = createSprites(&spritesheet, startX, startY+3*step, endX, endY+3*step, 16, 16)

	g.ghostSprites = ghosts
}

func (g *Game) loadSprites() {
	g.loadBackgroundSprites()
	g.loadGhostSprites()
	g.loadShotSprites()
	g.loadPacSprites()
	g.loadShipSprites()
}
