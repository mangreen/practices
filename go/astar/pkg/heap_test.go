package pkg

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Heap(t *testing.T) {
	assert := assert.New(t)

	h := &InitHeap{
		&Point{F: 2},
		&Point{F: 1},
		&Point{F: 3},
	}
	heap.Init(h)

	assert.Equal(3, len(*h), "Should be equal")

	heap.Push(h, &Point{F: 4})

	assert.Equal(4, len(*h), "Should be equal")

	F := 1
	for h.Len() > 0 {
		assert.Equal(F, heap.Pop(h).(*Point).F, "Should be equal")
		F++
	}

	assert.Equal(0, len(*h), "Should be equal")
}
