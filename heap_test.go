package heap

import (
	"math/rand"
	"testing"
)

type Int int

func (i Int) Less(j Node) bool {
	return i < j.(Int)
}

func TestHeap(t *testing.T) {
	h := New()
	randRange := 100
	numNodes := 100
	for i := 0; i < numNodes; i++ {
		h.Push(Int(rand.Intn(randRange)))
	}
}
