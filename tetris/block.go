package tetris

import (
	"math/rand"
)

type Block struct {
	p []Point
}

func NewBlock(points []Point) Block {
	return Block{p: points}

}

var blocks map[int]Block

func createBlocks() {
	blocks = map[int]Block{
		0: NewBlock([]Point{{0, 0}, {0, 1}, {1, 1}, {1, 2}}),
		1: NewBlock([]Point{{0, 0}, {0, 1}, {1, 0}, {1, 1}}),
	}
}

func (gm *Game) genRandomBlock() {
	r := rand.Intn(len(blocks))
	gm.addBlock(blocks[r])
}
