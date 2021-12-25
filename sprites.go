package main

import "github.com/faiface/pixel"

func createSprites(spritesheet *pixel.Picture, minX, minY, maxX, maxY, stepX, stepY float64) []*pixel.Sprite {
	sprites := []*pixel.Sprite{}

	for y := minY; y < maxY; y += stepY {
		for x := minX; x < maxX; x += stepX {
			frame := pixel.R(x, y, x+stepX, y+stepY)
			sprites = append(sprites, pixel.NewSprite(*spritesheet, frame))
		}
	}

	return sprites
}
