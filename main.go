package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	window, screen := setup()
	assets := loadAssets()

	assets.Board.Blit(nil, screen, nil)
	window.UpdateSurface()

	sdl.Delay(5000)

	exit(window, assets)
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

func loadAssets() *Board {
	board := &Board{}

	board.Board = sdl.LoadBMP("assets/grid.bmp")
	board.X = sdl.LoadBMP("assets/x.bmp")
	board.O = sdl.LoadBMP("assets/o.bmp")
	board.Validate()

	return board
}

func exit(window *sdl.Window, assets *Board) {
	assets.Free()
	window.Destroy()
	sdl.Quit()
}
