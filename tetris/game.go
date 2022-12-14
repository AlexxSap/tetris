package tetris

import (
	"sync"

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
	moveMutex      sync.Mutex
}

func newGame() *Game {
	rowCount := 15
	columnCount := 10

	createBlocks()

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

func Start() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	var gameOverChanel chan bool = make(chan bool)

	game := newGame()

	canvas.ClearScreen()
	game.drawBoxes()

	game.genRandomBlock()

	go game.listenKeyboard()
	go game.repaint()
	go game.move(gameOverChanel)

	<-gameOverChanel
	game.isOver = true

}
