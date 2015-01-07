package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

var blockCoords = [9][2]int{
	{1, 0}, {35, 0}, {68, 0},
	{1, 34}, {35, 34}, {68, 34},
	{1, 67}, {35, 67}, {68, 67},
}

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
	b.placement[block-1] = &Block{symbol, block}
}

func (b *Board) Draw(screen *sdl.Surface) {
	b.Board.Blit(nil, screen, nil)
	for i := 0; i < len(b.placement); i++ {
		if b.placement[i] == nil {
			continue
		}

		c := blockCoords[i]

		switch b.placement[i].Symbol {
		case SYMBOL_X:
			b.X.Blit(nil, screen, &sdl.Rect{int32(c[0]), int32(c[1]), 30, 30})
		case SYMBOL_O:
			b.O.Blit(nil, screen, &sdl.Rect{int32(c[0]), int32(c[1]), 30, 30})
		}
	}
}

const (
	SYMBOL_NONE = iota
	SYMBOL_X    = iota
	SYMBOL_O    = iota
)

type Block struct {
	Symbol int
	Dest   int
}

func (b *Block) Free() bool {
	return b.Symbol == SYMBOL_NONE
}
