package main

import (
	"math/rand"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Tree struct {
	x      float64
	y      float64
	sprite *pixel.Sprite
	matrix pixel.Matrix
}

type Forest struct {
	spritesheet     pixel.Picture
	availableFrames []pixel.Rect
	trees           []*Tree
}

func NewForest() *Forest {
	spritesheet, err := loadPicture("trees.png")
	if err != nil {
		panic(err)
	}

	frames := []pixel.Rect{}
	for x := spritesheet.Bounds().Min.X; x < spritesheet.Bounds().Max.X; x += 32 {
		for y := spritesheet.Bounds().Min.Y; y < spritesheet.Bounds().Max.Y; y += 32 {
			frames = append(frames, pixel.R(x, y, x+32, y+32))
		}
	}

	return &Forest{
		spritesheet:     spritesheet,
		availableFrames: frames,
		trees:           []*Tree{},
	}
}

func (f *Forest) AddTree(pos pixel.Vec) {
	tree := pixel.NewSprite(f.spritesheet, f.availableFrames[rand.Intn(len(f.availableFrames))])
	trees = append(trees, tree)
	treeMatrix := pixel.IM.Scaled(pixel.ZV, 1.5).Moved(pos)
	treeMatrices = append(treeMatrices, treeMatrix)
}

func (f *Forest) Draw(win *pixelgl.Window) {
	for i, tree := range trees {
		tree.Draw(win, treeMatrices[i])
	}
}
