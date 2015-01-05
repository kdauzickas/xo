package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Board struct {
	Board     *sdl.Surface
	X         *sdl.Surface
	O         *sdl.Surface
	placement [9]*Block
}

func (b *Board) Validate() {
	if b.X == nil || b.Board == nil || b.O == nil {
		panic("Failed to load asset")
	}
}

func (b *Board) Unload() {
	b.Board.Free()
	b.X.Free()
	b.O.Free()
}

func (b *Board) Free(block int) bool {
	return 0 < block && block < 10 && (b.placement[block-1] == nil || b.placement[block-1].Free())
}

func (b *Board) Place(block, symbol int) {
	newSymbol := Block(symbol)
	b.placement[block-1] = &newSymbol
}

const (
	SYMBOL_NONE = iota
	SYMBOL_X    = iota
	SYMBOL_O    = iota
)

type Block int

func (b *Block) Free() bool {
	return *b == SYMBOL_NONE
}
