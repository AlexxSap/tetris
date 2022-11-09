package tetris

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

func (gm *Game) repaint() {

}
