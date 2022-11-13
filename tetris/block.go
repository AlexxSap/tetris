package tetris

import (
	"math/rand"
)

type Block struct {
	d []struct{ l, c int }
}

func NewBlock(points []struct{ l, c int }) Block {
	return Block{d: points}

}

var blocks map[int]Block

func createBlocks() {
	blocks = map[int]Block{
		0: NewBlock([]struct{ l, c int }{{0, 0}, {0, 1}, {1, 1}, {1, 2}}),
		// 1: []canvas.Point{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
	}
}

func (gm *Game) genRandomBlock() {
	r := rand.Intn(len(blocks))
	gm.addNewBlock(blocks[r])
}
