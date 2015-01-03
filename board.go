package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Board struct {
	Board *sdl.Surface
	X     *sdl.Surface
	O     *sdl.Surface
}

func (b *Board) Validate() {
	if b.X == nil || b.Board == nil || b.O == nil {
		panic("Failed to load asset")
	}
}

func (b *Board) Free() {
	b.Board.Free()
	b.X.Free()
	b.O.Free()
}
