package tetris

import (
	canvas "github.com/AlexxSap/SiDCo"
)

type Block struct {
	p []Point
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
	colOffset := 5 /// TODO заменить это, чтобы блоки было посреди
	blocks = map[int]Block{
		0: NewBlock([]Point{{1, colOffset + 0}, {1, colOffset + 1}, {2, colOffset + 1}, {2, colOffset + 2}}),
		1: NewBlock([]Point{{1, colOffset + 0}, {1, colOffset + 1}, {2, colOffset + 0}, {2, colOffset + 1}}),
	}
}

func (gm *Game) genRandomBlock() {
	// i := rand.Intn(len(blocks))
	i := 0
	gm.block = blocks[i]
}

func (b *Block) moveDown() {
	for i := 0; i < len(b.p); i++ {
		b.p[i].Line++
	}
}
