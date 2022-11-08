package tetris

import (
	"github.com/AlexxSap/matrix"
)

type Game struct {
	rowCount    int
	columnCount int
	field       *matrix.Matrix[int]
}

func newGame() *Game {
	rowCount := 20
	columnCount := 10
	d := make([]int, rowCount*columnCount, rowCount*columnCount)

	m, err := matrix.NewMatrix(d, rowCount, columnCount)
	if err != nil {
		panic(err)
	}

	return &Game{
		rowCount:    rowCount,
		columnCount: columnCount,
		field:       m,
	}
}

func Start() {
	game := newGame()

	go game.repaint()
	go game.move()
}
