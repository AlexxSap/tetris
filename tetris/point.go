package tetris

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type Point struct {
	Line, Column int
}

type PointIterator struct {
	data  []Point
	index int
}

func (p *PointIterator) Begin() {
	p.index = 0
}

func (p *PointIterator) Next() bool {
	if len(p.data) == 0 {
		return false
	}
	if p.index < len(p.data) {
		p.index++
		return true
	}
	return false
}

func (p *PointIterator) First() int {
	return p.data[p.index-1].Line
}

func (p *PointIterator) Second() int {
	return p.data[p.index-1].Column
}
