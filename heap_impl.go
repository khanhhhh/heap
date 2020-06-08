package heap

import "fmt"

type heap struct {
	nodes []Node
}

func (h *heap) heapConsistent() {
	if Debug && len(h.nodes) > 1 {
		for i := 0; i < len(h.nodes); i++ {
			left, right := childrenOf(i)
			if (h.inHeap(left) && h.nodes[left].Less(h.nodes[i])) || (h.inHeap(right) && h.nodes[right].Less(h.nodes[i])) {
				fmt.Println(i, left, right)
				panic("heap inconsistent")
			}
		}
	}
}

func childrenOf(i int) (int, int) {
	idx := i + 1
	idx1 := 2 * idx
	idx2 := 2*idx + 1
	return idx1 - 1, idx2 - 1
}

func parentOf(i int) int {
	idx := i + 1
	idx0 := idx / 2
	return idx0 - 1
}

func (h *heap) isLeaf(i int) bool {
	child1, _ := childrenOf(i)
	return child1 >= len(h.nodes)
}

func (h *heap) inHeap(i int) bool {
	return i >= 0 && i < len(h.nodes)
}

func (h *heap) fixHeapTopDown(parent int) {
	if h.isLeaf(parent) == false {
		left, right := childrenOf(parent)
		var less int = left
		if h.inHeap(right) && h.nodes[right].Less(h.nodes[left]) {
			less = right
		}
		if h.nodes[less].Less(h.nodes[parent]) {
			h.nodes[parent], h.nodes[less] = h.nodes[less], h.nodes[parent]
			h.fixHeapTopDown(less)
		}
	}
}

func (h *heap) fixHeapBottomUp(child int) {
	if child > 0 {
		parent := parentOf(child)
		left, right := childrenOf(parent)
		var less int = left
		if h.inHeap(right) && h.nodes[right].Less(h.nodes[left]) {
			less = right
		}
		if h.nodes[less].Less(h.nodes[parent]) {
			h.nodes[parent], h.nodes[less] = h.nodes[less], h.nodes[parent]
			h.fixHeapBottomUp(parent)
		}
	}
}

func (h *heap) Len() int {
	return len(h.nodes)
}

func (h *heap) Push(node Node) {
	defer h.heapConsistent()
	h.nodes = append(h.nodes, node)
	h.fixHeapBottomUp(len(h.nodes) - 1)
}

func (h *heap) Pop() (node Node) {
	defer h.heapConsistent()
	if len(h.nodes) > 0 {
		node = h.nodes[0]
		h.nodes[0] = h.nodes[len(h.nodes)-1]
		h.nodes = h.nodes[:len(h.nodes)-1]
		h.fixHeapTopDown(0)
	} else {
		node = nil
	}
	return node
}
func (h *heap) heapify(i int) {
	if !h.isLeaf(i) {
		left, right := childrenOf(i)
		h.heapify(left)
		h.heapify(right)
		h.fixHeapTopDown(i)
	}
}
