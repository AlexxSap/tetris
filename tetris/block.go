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
		1: NewBlock([]Point{{0, 0}, {0, 1}, {1, 0}, {1, 1}}),
		2: NewBlock([]Point{{0, 0}, {1, 0}, {2, 0}, {3, 0}}),
		3: NewBlock([]Point{{0, 0}, {1, 0}, {2, 0}, {2, 1}}),
		0: NewBlock([]Point{{0, 0}, {0, 1}, {1, 1}, {1, 2}}),
		4: NewBlock([]Point{{0, 1}, {1, 1}, {1, 0}, {2, 1}}),
		5: NewBlock([]Point{{0, 1}, {1, 0}, {1, 1}, {2, 1}}),
		6: NewBlock([]Point{{0, 1}, {1, 1}, {2, 1}, {2, 0}}),
	}
}

func (gm *Game) genRandomBlock() {
	rand.Seed(int64(time.Now().Nanosecond()))
	gm.block = blocks[rand.Intn(len(blocks))]
	gm.currentStep++
	gm.moveRightBlock(gm.columnCount / 2)
}

func (gm *Game) addCurrentBlockToTheBottom() {
	for _, p := range gm.block.p {
		gm.field.Set(p.Line, p.Column, gm.currentStep)
	}
}

func (gm *Game) moveRightBlock(offset int) {
	for i := 0; i < len(gm.block.p); i++ {
		newVal := gm.block.p[i].Column + offset
		if newVal <= 0 || newVal > gm.columnCount+1 {
			return
		}
	}

	for i := 0; i < len(gm.block.p); i++ {
		gm.block.p[i].Column += offset
	}
}

func (gm *Game) moveDownBlock(offset int) {
	for i := 0; i < len(gm.block.p); i++ {
		gm.block.p[i].Line += offset
	}
}

func (b *Block) offsets() (int, int) {
	col, row := b.p[0].Column, b.p[0].Line
	for _, p := range b.p {
		col, row = min(col, p.Column), min(row, p.Line)
	}
	return col, row
}

func (b *Block) rotate() {
	x, y := b.offsets()

	moveRight := func(b *Block, val int) {
		for i := 0; i < len(b.p); i++ {
			b.p[i].Column += val
		}
	}

	moveDown := func(b *Block, val int) {
		for i := 0; i < len(b.p); i++ {
			b.p[i].Line += val
		}
	}

	moveRight(b, -x)
	moveDown(b, -y)
	m := matrix.NewMatrixFromPoints(b.iterator(), 666)
	m.Rotate()

	points, err := m.Filtered(func(cell int) bool { return cell == 666 })
	if err != nil {
		panic(err)
	}

	p := make([]Point, 0, len(points))
	for _, point := range points {
		p = append(p, Point{point.Row, point.Column})
	}

	/// TODO при повороте фигура у края может заехать на границу

	b.p = p

	moveRight(b, x)
	moveDown(b, y)
}
