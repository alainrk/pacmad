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

// var (
// 	camPos       = pixel.ZV
// 	camSpeed     = 500.0
// 	camZoom      = 1.0
// 	camZoomSpeed = 1.01
// )

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
	game := NewGame(win)
	lastTime := time.Now()

	// Main Loop
	for !win.Closed() {
		// Keep consistent FPS rate, adjusting the rotation (in this case) according to the time elapsed since the last frame.
		dt := time.Since(lastTime).Seconds()
		lastTime = time.Now()

		// -- Cam
		// cam := pixel.IM.Scaled(camPos, camZoom).Moved(win.Bounds().Center().Sub(camPos))
		// win.SetMatrix(cam)

		// if win.Pressed(pixelgl.KeyLeft) {
		// 	camPos.X -= camSpeed * dt
		// }
		// if win.Pressed(pixelgl.KeyRight) {
		// 	camPos.X += camSpeed * dt
		// }
		// if win.Pressed(pixelgl.KeyDown) {
		// 	camPos.Y -= camSpeed * dt
		// }
		// if win.Pressed(pixelgl.KeyUp) {
		// 	camPos.Y += camSpeed * dt
		// }
		// camZoom *= math.Pow(camZoomSpeed, win.MouseScroll().Y)

		// --- Actions
		pac.Move(-dt)
		game.Update()

		if win.JustPressed(pixelgl.MouseButtonLeft) {
			// mouse := cam.Unproject(win.MousePosition())
			mouse := win.MousePosition()
			game.AddShot(mouse)
		}

		// --- Draw
		win.Clear(colornames.Black)

		// Pac
		pac.Draw(win)

		// Trees
		game.Draw(win)

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
