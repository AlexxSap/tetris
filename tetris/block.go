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
	b := blocks[rand.Intn(len(blocks))]
	b.moveRight(gm.columnCount / 2)
	gm.block = b
}

func (b *Block) moveRight(offset int) {
	for i := 0; i < len(b.p); i++ {
		b.p[i].Column += offset
	}
}

func (b *Block) moveDown(offset int) {
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

func (b *Block) rotate() {
	x, y := b.offsets()
	b.moveRight(-x)
	b.moveDown(-y)
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
	b.p = p

	b.moveRight(x)
	b.moveDown(y)
}
