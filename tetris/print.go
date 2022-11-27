package tetris

import (
	"time"

	canvas "github.com/AlexxSap/SiDCo"
)

func colorByValue(value int) canvas.Color {
	return canvas.Color(value%8 + 1)
}

const (
	block = "\u2593"
)

func (gm *Game) drawBoxes() {
	gm.blocksField.DrawBoxWithTitle("")
	gm.nextBlockField.DrawBoxWithTitle("NEXT")
	gm.infoField.DrawBoxWithTitle("INFO")
}

func (gm *Game) repaintCurrentBlock() {
	points := gm.block.canvasPoints()
	if len(points) != 0 {
		gm.blocksField.DrawPath(colorByValue(gm.currentStep).String()+block, points)
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
