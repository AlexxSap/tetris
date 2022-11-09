package tetris

import (
	canvas "github.com/AlexxSap/SiDCo"
	"github.com/AlexxSap/matrix"
)

type Game struct {
	rowCount       int
	columnCount    int
	field          *matrix.Matrix[int]
	blocksField    canvas.Canvas
	nextBlockField canvas.Canvas
	infoField      canvas.Canvas
}

func newGame() *Game {
	rowCount := 15
	columnCount := 10
	d := make([]int, rowCount*columnCount)

	m, err := matrix.NewMatrix(d, rowCount, columnCount)
	if err != nil {
		panic(err)
	}

	blocksField, _ := canvas.NewCanvas(canvas.Point{Line: 1, Column: 10}, canvas.Point{Line: rowCount + 1, Column: columnCount + 1})
	nextBlockField, _ := canvas.NewCanvas(canvas.Point{Line: 1, Column: 25}, canvas.Point{Line: 5, Column: 10})
	infoField, _ := canvas.NewCanvas(canvas.Point{Line: 7, Column: 25}, canvas.Point{Line: 10, Column: 10})

	return &Game{
		rowCount:       rowCount,
		columnCount:    columnCount,
		field:          m,
		blocksField:    blocksField,
		nextBlockField: nextBlockField,
		infoField:      infoField,
	}
}

func Start() {
	game := newGame()

	canvas.ClearScreen()
	game.drawBoxes()

	// go game.repaint()
	// go game.move()
}
