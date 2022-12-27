package tetris

import (
	"math/rand"
	"time"

	canvas "github.com/AlexxSap/SiDCo"
	"github.com/AlexxSap/matrix"
)

type Block struct {
	p []Point
}

func (b *Block) iterator() *PointIterator {
	return &PointIterator{b.p, 0}
}

func (b *Block) canvasPoints() []canvas.Point {
	p := make([]canvas.Point, 0, len(b.p))
	for _, point := range b.p {
		p = append(p, canvas.Point{Line: point.Line, Column: point.Column})
	}
	return p
}

func NewBlock(points []Point) Block {
	return Block{p: points}
}

var blocks map[int]Block

func createBlocks() {
	blocks = map[int]Block{
		0: NewBlock([]Point{{0, 0}, {1, 0}, {2, 0}, {3, 0}}),
		1: NewBlock([]Point{{0, 0}, {0, 1}, {1, 0}, {1, 1}}),
		2: NewBlock([]Point{{0, 0}, {1, 0}, {2, 0}, {2, 1}}),
		3: NewBlock([]Point{{0, 0}, {0, 1}, {1, 1}, {1, 2}}),
		4: NewBlock([]Point{{0, 1}, {1, 1}, {1, 0}, {2, 1}}),
		5: NewBlock([]Point{{0, 1}, {1, 0}, {1, 1}, {2, 1}}),
		6: NewBlock([]Point{{0, 1}, {1, 1}, {2, 1}, {2, 0}}),
	}
}

func (gm *Game) genRandomBlock() {
	rand.Seed(int64(time.Now().Nanosecond()))

	if gm.nextBlock == -1 {
		gm.nextBlock = rand.Intn(len(blocks))
	}

	block := blocks[gm.nextBlock]
	gm.nextBlock = rand.Intn(len(blocks))
	gm.block.p = make([]Point, len(block.p))
	copy(gm.block.p, block.p)
	gm.currentStep++
	gm.block.moveDownBy(1)
	gm.block.moveRightBy(gm.columnCount / 2)
	gm.printNexBlock()
}

func (gm *Game) addCurrentBlockToTheBottom() {
	for _, p := range gm.block.p {
		gm.field.Set(p.Line, p.Column, gm.currentStep)
	}
}

func (gm *Game) rowsToDestroy() []int {
	res := make([]int, 0)
	for row := 0; row <= gm.rowCount; row++ {
		zeroFounded := false
		for col := 1; col <= gm.columnCount; col++ {
			val, _ := gm.field.Get(row, col)
			if val == 0 {
				zeroFounded = true
				break
			}
		}
		if !zeroFounded {
			res = append(res, row)
		}
	}
	return res
}

func (gm *Game) destroyRows(rows []int) {
	for _, row := range rows {
		err := gm.field.RemoveRow(row)
		if err != nil {
			panic(err)
		}
		time.Sleep(100 * time.Millisecond)
		gm.repaintAllBlocks()
	}
}

func (gm *Game) canMoveRight(offset int) bool {
	for i := 0; i < len(gm.block.p); i++ {
		newVal := gm.block.p[i].Column + offset
		if newVal <= 0 || newVal > gm.columnCount {
			return false
		}
	}

	gm.block.moveRightBy(offset)
	match, _ := gm.field.AnyOfPoints(
		gm.block.iterator(),
		func(val int) bool {
			return val > 0
		})

	gm.block.moveRightBy(-offset)

	return !match
}

func (gm *Game) moveBlockLeft() {
	if gm.canMoveRight(-1) {
		gm.block.moveRightBy(-1)
	}
}

func (gm *Game) moveBlockRight() {
	if gm.canMoveRight(1) {
		gm.block.moveRightBy(1)
	}
}

func (b *Block) moveRightBy(offset int) {
	for i := 0; i < len(b.p); i++ {
		b.p[i].Column += offset
	}
}

func (b *Block) moveDownBy(offset int) {
	for i := 0; i < len(b.p); i++ {
		b.p[i].Line += offset
	}
}

func (b *Block) offsets() (int, int) {
	col, row := b.p[0].Column, b.p[0].Line
	for _, p := range b.p {
		col, row = min(col, p.Column), min(row, p.Line)
	}
	return col, row
}

func (gm *Game) rotate() {

	var tempBlock Block
	tempBlock.p = make([]Point, len(gm.block.p))
	copy(tempBlock.p, gm.block.p)

	x, y := tempBlock.offsets()

	tempBlock.moveRightBy(-x)
	tempBlock.moveDownBy(-y)
	m := matrix.NewMatrixFromPoints(tempBlock.iterator(), 666)
	m.Rotate()

	points, err := m.Filtered(func(cell int) bool { return cell == 666 })
	if err != nil {
		panic(err)
	}

	p := make([]Point, 0, len(points))
	for _, point := range points {
		if point.Column+x > gm.columnCount {
			return
		}

		p = append(p, Point{point.Row, point.Column})
	}

	copy(tempBlock.p, p)
	tempBlock.moveRightBy(x)
	tempBlock.moveDownBy(y)

	copy(gm.block.p, tempBlock.p)

}
