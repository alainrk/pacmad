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
	game := NewGame(win)

	// Main Loop
	for !win.Closed() {
		// --- Actions
		game.Update()
		if win.JustReleased(pixelgl.KeySpace) || win.JustPressed(pixelgl.MouseButtonLeft) {
			if game.status == "gameover" {
				game = NewGame(win)
			} else {
				game.AddShot(win.Bounds().Center(), win.MousePosition())
			}
		}

		// --- Draw
		win.Clear(colornames.Black)

		// Game
		game.Draw()

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
	rand.Seed(time.Now().UnixNano())
	// Locks current goroutine on it's current thread
	// https://github.com/faiface/pixel/wiki/Creating-a-Window#run
	pixelgl.Run(run)
	// Calling pixelgl.Run puts PixelGL in control of the main function.
	// There's no way for us to run any code in the main function anymore.
}
