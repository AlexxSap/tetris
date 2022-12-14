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
		case keyboard.KeyArrowLeft:
			gm.moveBlockLeft()
		case keyboard.KeyArrowRight:
			gm.moveBlockRight()
		case keyboard.KeyArrowUp:
			gm.rotate()
		case keyboard.KeyEsc:
			gm.isOver = true
		}
		gm.drawCurrentBlock()
	}
}

func (gm *Game) addToTheBottom() bool {
	gm.addCurrentBlockToTheBottom()

	if rows := gm.rowsToDestroy(); len(rows) != 0 {
		time.Sleep(100 * time.Millisecond)
		gm.destroyRows(rows)
	}
	time.Sleep(100 * time.Millisecond)
	gm.printInfo()
	gm.genRandomBlock()
	return !gm.isCurrentBlockAtTheBottom()
}

func (gm *Game) isCurrentBlockAtTheBottom() bool {
	for _, p := range gm.block.p {
		if p.Line == gm.rowCount {
			return true
		}
	}

	/// проверяем, что не наткнулись на другую фигуру
	gm.block.moveDownBy(1)
	match, _ := gm.field.AnyOfPoints(
		gm.block.iterator(),
		func(val int) bool {
			return val > 0
		})
	gm.block.moveDownBy(-1)

	return match
}

func (gm *Game) slideDown() {
	gm.moveMutex.Lock()
	defer gm.moveMutex.Unlock()

	gm.clearCurrentBlock()
	for !gm.isCurrentBlockAtTheBottom() {
		gm.block.moveDownBy(1)
	}
	gm.drawCurrentBlock()

	if !gm.addToTheBottom() {
		gm.isOver = true
	}
}

func (gm *Game) move(gameOverChanel chan<- bool) {

	defer gm.moveMutex.Unlock()
	for {
		gm.moveMutex.Lock()

		if gm.isOver {
			gameOverChanel <- true
			return
		}

		gm.clearCurrentBlock()
		gm.block.moveDownBy(1)
		gm.drawCurrentBlock()
		if gm.isCurrentBlockAtTheBottom() {
			if !gm.addToTheBottom() {
				gameOverChanel <- true
				return
			}
		}

		gm.moveMutex.Unlock()
		time.Sleep(1 * time.Second)
	}

}
