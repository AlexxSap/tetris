package tetris

import (
	canvas "github.com/AlexxSap/SiDCo"
)

func colorByValue(value int) canvas.Color {
	return canvas.Color(value%8 + 1)
}

const (
	block = "\u2593"
	empty = " "
)

func (gm *Game) drawBoxes() {
	gm.blocksField.DrawBoxWithTitle("")
	gm.nextBlockField.DrawBoxWithTitle("NEXT")
	gm.infoField.DrawBoxWithTitle("INFO")
}

func (gm *Game) clearCurrentBlock() {
	points := gm.block.canvasPoints()
	if len(points) != 0 {
		gm.blocksField.DrawPath(empty, points)
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
