package main

import (
	"image"
	"math/rand"
	"os"
	"time"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func RandIntInRange(min, max int) int {
	return min + rand.Intn(max-min)
}

var (
	trees        []*pixel.Sprite
	treeMatrices []pixel.Matrix
)

type Pac struct {
	x      float64
	y      float64
	angle  float64
	sprite *pixel.Sprite
}

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

func (p *Pac) move(dx, dy, maxx, maxy float64) {
	p.x += dx
	p.y += dy
}

func NewPac() *Pac {
	pic, err := loadPicture("pac.png")
	if err != nil {
		panic(err)
	}

	sprite := pixel.NewSprite(pic, pic.Bounds())

	return &Pac{
		x:      0,
		y:      0,
		angle:  0,
		sprite: sprite,
	}
}

func CreateWindow() *pixelgl.Window {
	cfg := pixelgl.WindowConfig{
		Title:  "Pacmad",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true, // Window updates only as often as the monitor refreshes, and not as often as the loop does.
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	// win.SetSmooth(true)
	return win
}

func run() {
	win := CreateWindow()
	pac := NewPac()
	forest := NewForest()
	lastTime := time.Now()

	// Main Loop
	for !win.Closed() {
		// Keep consistent FPS rate, adjusting the rotation (in this case) according to the time elapsed since the last frame.
		dt := time.Since(lastTime).Seconds()
		lastTime = time.Now()

		// --- Actions
		pac.angle += 5 * dt
		pacmove := 1000 * dt
		pac.move(pacmove, pacmove, win.Bounds().W(), win.Bounds().H())

		if win.JustPressed(pixelgl.MouseButtonLeft) {
			forest.AddTree(win.MousePosition())
		}

		// --- Draw
		win.Clear(colornames.Black)

		// Pac
		mat := pixel.IM
		mat = mat.Moved(pixel.V(pac.x, pac.y))
		mat = mat.ScaledXY(win.Bounds().Center(), pixel.V(0.02, 0.02))
		// mat = mat.Rotated(win.Bounds().Center(), pac.angle)
		pac.sprite.Draw(win, mat)

		// Trees
		for i, tree := range trees {
			tree.Draw(win, treeMatrices[i])
		}

		win.Update()
	}
}

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func main() {
	// Locks current goroutine on it's current thread
	// https://github.com/faiface/pixel/wiki/Creating-a-Window#run
	pixelgl.Run(run)
	// Calling pixelgl.Run puts PixelGL in control of the main function.
	// There's no way for us to run any code in the main function anymore.
}
