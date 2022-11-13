package tetris

import (
	canvas "github.com/AlexxSap/SiDCo"
	"github.com/AlexxSap/matrix"
	"github.com/eiannone/keyboard"
)

type Game struct {
	rowCount       int
	columnCount    int
	block          Block
	field          *matrix.Matrix[int]
	blocksField    canvas.Canvas
	nextBlockField canvas.Canvas
	infoField      canvas.Canvas
	isOver         bool
	currentStep    int
}

func newGame() *Game {
	createBlocks()
	rowCount := 15
	columnCount := 10

	m, err := matrix.NewMatrix(make([]int, rowCount*columnCount),
		rowCount,
		columnCount)
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
		isOver:         false,
		currentStep:    1,
	}
}

func (gm *Game) addBlock(block Block) {
	for _, p := range block.p {
		gm.field.Set(p.Line, p.Column, gm.currentStep)
	}
}

func (gm *Game) remBlock(block Block) {
	for _, p := range block.p {
		gm.field.Set(p.Line, p.Column, 0)
	}
}

func Start() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	game := newGame()

	canvas.ClearScreen()
	game.drawBoxes()

	go game.repaint()
	go game.move()
}
