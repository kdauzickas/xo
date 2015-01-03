package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

var grid *sdl.Surface

func main() {
	window, screen := setup()
	loadAssets()

	grid.Blit(nil, screen, nil)
	window.UpdateSurface()

	sdl.Delay(5000)

	exit(window)
}

func setup() (*sdl.Window, *sdl.Surface) {
	if sdl.Init(sdl.INIT_EVERYTHING) < 0 {
		panic("Couldn't init SDL")
	}
	window, err := sdl.CreateWindow(
		"Tic-tac-toe",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		100,
		100,
		sdl.WINDOW_SHOWN,
	)

	if err != nil {
		panic(err)
	}

	return window, window.GetSurface()
}

func loadAssets() {
	grid = sdl.LoadBMP("assets/grid.bmp")
	if grid == nil {
		panic("Failed to load grid")
	}
}

func exit(window *sdl.Window) {
	grid.Free()
	window.Destroy()
	sdl.Quit()
}
