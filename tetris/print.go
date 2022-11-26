package tetris

import (
	"time"
)

// import (
// 	canvas "github.com/AlexxSap/SiDCo"
// )

const (
	block = "\u001b[31;1m\u2593" /// TODO \u001b[31;1m - цвет
)

func (gm *Game) drawBoxes() {
	gm.blocksField.DrawBoxWithTitle("")
	gm.nextBlockField.DrawBoxWithTitle("NEXT")
	gm.infoField.DrawBoxWithTitle("INFO")
}

func (gm *Game) repaintCurrentBlock() {
	points := gm.block.canvasPoints()
	// fmt.Println(points)
	if len(points) != 0 {
		gm.blocksField.DrawPath(block, points)
	}
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

		/// TODO delele this
		gm.blocksField.ClearInner()
		/// перерисовывать только текущий блок
		gm.repaintCurrentBlock()
		if gm.needRepaintAllBlocks() {
			time.Sleep(time.Duration(repaintTime/3) * time.Millisecond)
			gm.repaintAllBlocks()
		}

		if gm.isOver {
			break
		}

		<-ticker.C
		resetTiker()
	}
}
