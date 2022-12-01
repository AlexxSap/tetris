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

func createBlocks(width int) {
	offset := width / 2
	blocks = map[int]Block{
		1: NewBlock([]Point{{0, offset + 0}, {0, offset + 1}, {1, offset + 0}, {1, offset + 1}}),
		2: NewBlock([]Point{{0, offset + 0}, {1, offset + 0}, {2, offset + 0}, {3, offset + 0}}),
		3: NewBlock([]Point{{0, offset + 0}, {1, offset + 0}, {2, offset + 0}, {2, offset + 1}}),
		0: NewBlock([]Point{{0, offset + 0}, {0, offset + 1}, {1, offset + 1}, {1, offset + 2}}),
		4: NewBlock([]Point{{0, offset + 1}, {1, offset + 1}, {1, offset + 0}, {2, offset + 1}}),
		5: NewBlock([]Point{{0, offset + 1}, {1, offset + 0}, {1, offset + 1}, {2, offset + 1}}),
		6: NewBlock([]Point{{0, offset + 1}, {1, offset + 1}, {2, offset + 1}, {2, offset + 0}}),
	}
}

func (gm *Game) genRandomBlock() {
	rand.Seed(int64(time.Now().Nanosecond()))
	gm.block = blocks[rand.Intn(len(blocks))]
}

func (b *Block) moveDown() {
	for i := 0; i < len(b.p); i++ {
		b.p[i].Line++
	}
}

func (b *Block) rotate() {
	m := matrix.NewSquareMatrixFromPoints(b.iterator(), 666)
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
}
