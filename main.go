package main

import (
	"image"
	"os"
	"time"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type Pac struct {
	x      float64
	y      float64
	angle  float64
	sprite *pixel.Sprite
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

	win.SetSmooth(true)
	return win
}

func run() {
	win := CreateWindow()
	pac := NewPac()
	lastTime := time.Now()

	// Main Loop
	for !win.Closed() {
		// Keep consistent FPS rate, adjusting the rotation (in this case) according to the time elapsed since the last frame.
		dt := time.Since(lastTime).Seconds()
		lastTime = time.Now()

		pac.angle += 5 * dt

		win.Clear(colornames.Black)

		mat := pixel.IM
		mat = mat.Moved(win.Bounds().Center())
		mat = mat.Rotated(win.Bounds().Center(), pac.angle)
		mat = mat.ScaledXY(win.Bounds().Center(), pixel.V(0.02, 0.02))

		pac.sprite.Draw(win, mat)

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
