package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_OpenList(t *testing.T) {
	assert := assert.New(t)

	open := new(OpenList)
	open.Init()

	open.Push(&Point{F: 2})
	open.Push(&Point{F: 1})
	open.Push(&Point{F: 3})

	assert.Equal(3, len(*open.Slice), "Should be equal")

	F := 1
	for len(*open.Slice) > 0 {
		assert.Equal(F, open.Pop().F, "Should be equal")
		F++
	}

	assert.Equal(0, len(*open.Slice), "Should be equal")
}
