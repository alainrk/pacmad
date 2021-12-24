package main

import "github.com/faiface/pixel"

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
	spritesheet, err := loadPicture("assets/shot.png")
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
	spritesheet, err := loadPicture("assets/pac.png")
	if err != nil {
		panic(err)
	}

	startX := spritesheet.Bounds().Min.X
	startY := spritesheet.Bounds().Min.Y
	endX := spritesheet.Bounds().Max.X
	endY := spritesheet.Bounds().Max.Y

	f.pacSprites = createSprites(&spritesheet, startX, startY, endX, endY, 32)
}

func (f *Game) loadShipSprites() {
	spritesheet, err := loadPicture("assets/ship.png")
	if err != nil {
		panic(err)
	}

	startX := spritesheet.Bounds().Min.X
	startY := spritesheet.Bounds().Min.Y
	endX := spritesheet.Bounds().Max.X
	endY := spritesheet.Bounds().Max.Y

	f.shipSprites = createSprites(&spritesheet, startX, startY, endX, endY, 32)
}

func (f *Game) loadGhostSprites() {
	spritesheet, err := loadPicture("assets/pmsprites.png")
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
	f.loadShipSprites()
}
