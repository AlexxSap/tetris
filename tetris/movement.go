package tetris

import "time"

func (gm *Game) move(gameOverChanel chan<- bool) {

	for i := 0; i < 5; i++ {
		gm.clearCurrentBlock()
		gm.block.moveDown(1)
		gm.block.rotate()
		gm.drawCurrentBlock()
		time.Sleep(1 * time.Second)
	}
	gameOverChanel <- true
}

func (gm *Game) needRepaintAllBlocks() bool {
	return false
}
