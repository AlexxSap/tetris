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

func (gm *Game) clearCurrentBlock() {
	/// TODO стирается верхушка поля
	points := gm.block.canvasPoints()
	if len(points) != 0 {
		gm.blocksField.DrawPath(" ", points)
	}
}

func (gm *Game) drawCurrentBlock() {
	points := gm.block.canvasPoints()
	if len(points) != 0 {
		gm.blocksField.SetColor(colorByValue(gm.currentStep))
		gm.blocksField.DrawPath(block, points)
		gm.blocksField.SetDefaultColor()
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
