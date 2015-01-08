package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	window, screen := setup()
	board := loadAssets()

	quit := false

	for !quit {
		won, symbol := board.HasWinner()

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				quit = true
			case *sdl.MouseButtonEvent:
				if !won {
					handleMouseClick(t, board)
				}
			}
		}

		if won {
			board.DrawWinner(screen, symbol)
		} else {
			board.Draw(screen)
		}

		window.UpdateSurface()
	}

	exit(window, board)
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
	assets.Unload()
	window.Destroy()
	sdl.Quit()
}

func handleMouseClick(click *sdl.MouseButtonEvent, board *Board) {
	if click.Type != sdl.MOUSEBUTTONUP {
		return
	}

	block := getBlock(click)

	if board.Free(block) {
		switch click.Button {
		case sdl.BUTTON_LEFT:
			board.Place(block, SYMBOL_X)
		case sdl.BUTTON_RIGHT:
			board.Place(block, SYMBOL_O)
		}
	}

	fmt.Println(board)
}

// Get block where the click happened
// Grid is numbered like this:
// 1 2 3
// 4 5 6
// 7 8 9
func getBlock(click *sdl.MouseButtonEvent) int {
	col := 0
	row := 0

	if click.X > 33 {
		col = 1
	}

	if click.X > 66 {
		col = 2
	}

	if click.Y > 33 {
		row = 1
	}

	if click.Y > 66 {
		row = 2
	}

	return row*3 + col + 1
}
