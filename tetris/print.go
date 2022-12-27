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

func (gm *Game) printNexBlock() {
	if gm.nextBlock == -1 {
		return
	}

	var b Block
	b.p = make([]Point, len(blocks[gm.nextBlock].p))
	copy(b.p, blocks[gm.nextBlock].p)
	b.moveDownBy(1)
	b.moveRightBy(gm.nextBlockField.Size().Column / 2)

	gm.nextBlockField.ClearInner()
	gm.nextBlockField.SetColor(colorByValue(gm.currentStep + 1))
	gm.nextBlockField.DrawPath(block, b.canvasPoints())
	gm.nextBlockField.SetDefaultColor()
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
	gm.blocksField.ClearInner()

	for row := 0; row <= gm.rowCount; row++ {
		for col := 1; col <= gm.columnCount; col++ {
			val, _ := gm.field.Get(row, col)
			if val > 0 {
				gm.blocksField.DrawColoredText(block, canvas.Point{Line: row, Column: col}, colorByValue(val))
			}
		}
	}
}
