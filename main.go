package main

import "github.com/veandco/go-sdl2/sdl"

func main() {
	window, err := sdl.CreateWindow("Tic-tac-toe", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		100, 100, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	surface := window.GetSurface()

	grid := sdl.LoadBMP("assets/grid.bmp")
	if grid == nil {
		panic("Failed to load grid")
	}

	grid.Blit(nil, surface, nil)

	window.UpdateSurface()

	sdl.Delay(5000)
	window.Destroy()
}
