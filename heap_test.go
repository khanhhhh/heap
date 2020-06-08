package heap

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type Int int

func (i Int) Less(j Node) bool {
	return i < j.(Int)
}

func timer(function func()) {
	t := time.Now()
	function()
	fmt.Printf("elapsed time: %v\n", time.Since(t))
}

func TestHeap(t *testing.T) {
	numHeaps := 1000
	numNodes := 1000
	randRange := 1000
	nodes := make([]Node, numNodes)
	for i := 0; i < numNodes; i++ {
		nodes[i] = Int(rand.Intn(randRange))
	}

	timer(func() {
		for t := 0; t < numHeaps; t++ {
			h := New()
			for _, node := range nodes {
				h.Push(node)
			}
		}
	})

	nodess := make([][]Node, numHeaps)
	for i := range nodess {
		nodess[i] = make([]Node, numNodes)
		for j := range nodess[i] {
			nodess[i][j] = nodes[j]
		}
	}
	timer(func() {
		for t := 0; t < numHeaps; t++ {
			h := FromArray(nodess[t])
			_ = h
		}
	})
}
