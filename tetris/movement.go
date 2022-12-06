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
			gm.slideDown()
			gm.addToTheBottom()
		case keyboard.KeyArrowLeft:
			gm.block.moveRight(-1, gm.columnCount)

		case keyboard.KeyArrowRight:
			gm.block.moveRight(1, gm.columnCount)
		case keyboard.KeyArrowUp:
			gm.block.rotate()
		}
		gm.drawCurrentBlock()
	}
}

/// TODO добавить в матрицу удаление строк со здвигом

func (gm *Game) addToTheBottom() bool {
	/// TODO добавить тут задержки
	// gm.addCurrentBlockToTheBottom()
	// if rows := gm.rowsToDestroy(); len(rows) != 0 {
	// 	gm.destroyRows(rows)
	// }
	// gm.genRandomBlock()
	return !gm.isCurrentBlockAtTheBottom()
}

func (gm *Game) isCurrentBlockAtTheBottom() bool {
	gm.block.moveDown(1)
	match, _ := gm.field.AnyOfPoints(
		gm.block.iterator(),
		func(val int) bool {
			return val > 0
		})
	gm.block.moveDown(-1)

	if match {
		return true
	}

	for _, p := range gm.block.p {
		if p.Line == gm.rowCount-1 {
			return true
		}
	}

	return false
}

func (gm *Game) slideDown() {
	for !gm.isCurrentBlockAtTheBottom() {
		gm.block.moveDown(1)
		time.Sleep(250 * time.Millisecond)
	}
}

func (gm *Game) move(gameOverChanel chan<- bool) {

	/// TODO del loop
	for i := 0; i < 5; i++ {
		gm.clearCurrentBlock()
		gm.block.moveDown(1)
		gm.drawCurrentBlock()
		if gm.isCurrentBlockAtTheBottom() {
			if !gm.addToTheBottom() {
				gameOverChanel <- true
				return
			}
		} else {
			time.Sleep(1 * time.Second)
		}
	}

	/// TODO cantAddCurrentBlock
	// if gm.cantAddCurrentBlock() {
	gameOverChanel <- true
	// }

}
