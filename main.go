package main

import "github.com/veandco/go-sdl2/sdl"

func main() {
	window, err := sdl.CreateWindow("Tic-tac-toe", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		600, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	surface := window.GetSurface()

	rect := sdl.Rect{0, 0, 200, 200}
	surface.FillRect(&rect, 0xffff0000)
	window.UpdateSurface()

	sdl.Delay(5000)
	window.Destroy()
}