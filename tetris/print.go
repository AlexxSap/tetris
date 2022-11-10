package tetris

import "time"

// import (
// 	canvas "github.com/AlexxSap/SiDCo"
// )

const (
	block = "\u2585"
)

func (gm *Game) drawBoxes() {
	gm.blocksField.DrawBoxWithTitle("")
	gm.nextBlockField.DrawBoxWithTitle("NEXT")
	gm.infoField.DrawBoxWithTitle("INFO")
}

func (gm *Game) repaintCurrentBlock() {

}

func (gm *Game) repaintAllBlocks() {

}

func (gm *Game) repaint() {
	repaintTime := 200

	var ticker *time.Timer
	resetTiker := func() {
		ticker = time.NewTimer(time.Millisecond * time.Duration(repaintTime))
	}

	resetTiker()

	for {
		/// перерисовывать только текущий блок
		gm.repaintCurrentBlock()
		time.Sleep(time.Duration(repaintTime/3) * time.Millisecond)
		if gm.needRepaintAllBlocks() {
			gm.repaintAllBlocks()
		}

		if gm.isOver {
			break
		}

		<-ticker.C
		resetTiker()
	}
}
