package tetris

import (
	"time"

	"github.com/eiannone/keyboard"
)

func (gm *Game) listenKeyboard() {
	for {
		if gm.isOver {
			break
		}

		_, key, err := keyboard.GetKey()
		if err != nil {
			break
		}

		gm.clearCurrentBlock()

		switch key {
		case keyboard.KeyArrowDown:
			/// TODO падение блока к остальным
			gm.slideDown()
			gm.addToTheBottom()
		case keyboard.KeyArrowLeft:
			gm.block.moveRight(-1)
		case keyboard.KeyArrowRight:
			gm.block.moveRight(1)
		case keyboard.KeyArrowUp:
			gm.block.rotate()
		}
		gm.drawCurrentBlock()
	}
}

/// TODO добавить в матрицу удаление строк со здвигом

func (gm *Game) addToTheBottom() {
	/// TODO добавить тут задержки
	gm.addCurrentBlockToTheBottom()
	if rows := gm.rowsToDestroy(); len(rows) != 0 {
		gm.destroyRows(rows)
	}
	gm.genRandomBlock()
}

func (gm *Game) move(gameOverChanel chan<- bool) {

	for i := 0; i < 5; i++ {
		gm.clearCurrentBlock()
		gm.block.moveDown(1)
		gm.drawCurrentBlock()
		if gm.isCurrentBlockAtTheBottom() {
			gm.addToTheBottom()
		} else {
			time.Sleep(1 * time.Second)
		}
	}
	gameOverChanel <- true
}
