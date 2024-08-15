package board

import (
	"periph.io/x/host/v3/orangepi"
)

type board struct {
}

func NewBoard() *board {
	return &board{}
}

func (b *board) MicName() string {
	return orangepi.FUN1_12.Name()
}
