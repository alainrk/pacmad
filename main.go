package main

import (
	"image"
	"os"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var i float64 = 0

func run() {
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

	pic, err := loadPicture("pacmad.png")
	if err != nil {
		panic(err)
	}

	sprite := pixel.NewSprite(pic, pic.Bounds())

	// Main Loop
	for !win.Closed() {
		i += 0.1

		win.Clear(colornames.Black)

		mat := pixel.IM
		mat = mat.Moved(win.Bounds().Center())
		mat = mat.Rotated(win.Bounds().Center(), i)
		mat = mat.ScaledXY(win.Bounds().Center(), pixel.V(0.02, 0.02))

		sprite.Draw(win, mat)

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
