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

		switch key {
		case keyboard.KeyArrowDown:
			gm.clearCurrentBlock()
			/// TODO падение блока к остальным
			// gm.block.moveDown(4)
			gm.drawCurrentBlock()
		case keyboard.KeyArrowLeft:
			gm.clearCurrentBlock()
			gm.block.moveRight(-1)
			gm.drawCurrentBlock()
		case keyboard.KeyArrowRight:
			gm.clearCurrentBlock()
			gm.block.moveRight(1)
			gm.drawCurrentBlock()
		case keyboard.KeyArrowUp:
			gm.clearCurrentBlock()
			/// TODO блок смещается вправо
			gm.block.rotate()
			gm.drawCurrentBlock()
		}
	}
}

func (gm *Game) move(gameOverChanel chan<- bool) {

	for i := 0; i < 5; i++ {
		gm.clearCurrentBlock()
		gm.block.moveDown(1)
		gm.drawCurrentBlock()
		time.Sleep(1 * time.Second)
	}
	gameOverChanel <- true
}

func (gm *Game) needRepaintAllBlocks() bool {
	return false
}
