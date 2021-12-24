package main

func (p *Panel) loadHearthSprites() {
	spritesheet, err := loadPicture("assets/hearth.png")
	if err != nil {
		panic(err)
	}

	step := 32.0
	startX := 0.0
	startY := 0.0
	endX := spritesheet.Bounds().Max.X
	endY := spritesheet.Bounds().Max.Y

	p.hearthSprites = createSprites(&spritesheet, startX, startY, endX, endY, step)
}
