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
			gm.moveRightBlock(-1)
		case keyboard.KeyArrowRight:
			gm.moveRightBlock(1)
		case keyboard.KeyArrowUp:
			gm.block.rotate()
		}
		gm.drawCurrentBlock()
	}
}

/// TODO добавить в матрицу удаление строк со здвигом

func (gm *Game) addToTheBottom() bool {
	/// TODO добавить тут задержки
	gm.addCurrentBlockToTheBottom()
	if rows := gm.rowsToDestroy(); len(rows) != 0 {
		gm.destroyRows(rows)
	}
	gm.genRandomBlock()
	return !gm.isCurrentBlockAtTheBottom()
}

func (gm *Game) isCurrentBlockAtTheBottom() bool {
	gm.moveDownBlock(1)
	match, _ := gm.field.AnyOfPoints(
		gm.block.iterator(),
		func(val int) bool {
			return val > 0
		})
	gm.moveDownBlock(-1)

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
		gm.moveDownBlock(1)
		time.Sleep(250 * time.Millisecond)
	}
}

/// TODO cantAddCurrentBlock
func (gm *Game) cantAddCurrentBlock() bool {
	return false
}

func (gm *Game) move(gameOverChanel chan<- bool) {

	for {
		gm.clearCurrentBlock()
		gm.moveDownBlock(1)
		gm.drawCurrentBlock()
		if gm.isCurrentBlockAtTheBottom() {
			if !gm.addToTheBottom() {
				gameOverChanel <- true
				return
			}
		} else {
			time.Sleep(1 * time.Second)
		}

		if gm.cantAddCurrentBlock() {
			gameOverChanel <- true
			break
		}
	}

}
