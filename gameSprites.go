package main

import "fmt"

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

	step := 16.0
	startX := spritesheet.Bounds().Max.X/3*2 + 3
	startY := 168.0
	endX := startX + (step * 8)
	endY := startY + step

	g.ghostSprites = createSprites(&spritesheet, startX, startY, endX, endY, step, step)
}

func (g *Game) loadSprites() {
	g.loadBackgroundSprites()
	g.loadGhostSprites()
	g.loadShotSprites()
	g.loadPacSprites()
	g.loadShipSprites()
}
