package pkg

import (
	"container/heap"
)

type OpenList struct {
	Slice *InitHeap
}

func (ol *OpenList) Init() {
	ol.Slice = &InitHeap{}
	heap.Init(ol.Slice)
}

func (ol *OpenList) Push(p *Point) {
	heap.Push(ol.Slice, p)
}

func (ol *OpenList) Pop() *Point {
	return heap.Pop(ol.Slice).(*Point)
}

func (ol *OpenList) IndexOf(x, y int) *Point {
	for _, p := range *ol.Slice {
		if p.X == x && p.Y == y {
			return p
		}
	}

	return nil
}