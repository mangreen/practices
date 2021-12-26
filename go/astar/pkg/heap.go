package pkg

type InitHeap []*Point

func (h InitHeap) Len() int { return len(h) }

func (h InitHeap) Less(i, j int) bool {
	return h[i].F < h[j].F
}

func (h InitHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *InitHeap) Pop() interface{} {
	old := *h
	l := len(old)
	p := old[l-1]
	*h = old[0 : l-1]
	return p
}

func (h *InitHeap) Push(x interface{}) {
	p := x.(*Point)
	*h = append(*h, p)
}