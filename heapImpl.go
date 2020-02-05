package heap

type heap struct {
	array []Node
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
	return child1 >= len(h.array)
}

func (h *heap) inHeap(i int) bool {
	return i >= 0 && i < len(h.array)
}

func (h *heap) fixHeapTopDown(parent int) {
	if h.isLeaf(parent) == false {
		left, right := childrenOf(parent)
		var less int = left
		if h.inHeap(right) && h.array[right].Less(h.array[left]) {
			less = right
		}
		if h.array[less].Less(h.array[parent]) {
			h.array[parent], h.array[less] = h.array[less], h.array[parent]
			h.fixHeapTopDown(less)
		}
	}
}

func (h *heap) fixHeapBottomUp(child int) {
	if child > 0 {
		parent := parentOf(child)
		left, right := childrenOf(parent)
		var less int = left
		if h.inHeap(right) && h.array[right].Less(h.array[left]) {
			less = right
		}
		if h.array[less].Less(h.array[parent]) {
			h.array[parent], h.array[less] = h.array[less], h.array[parent]
			h.fixHeapBottomUp(parent)
		}
	}
}

func (h *heap) Len() int {
	return len(h.array)
}

func (h *heap) Push(node Node) {
	h.array = append(h.array, node)
	h.fixHeapBottomUp(len(h.array) - 1)
}

func (h *heap) Pop() (node Node) {
	if len(h.array) > 0 {
		node = h.array[0]
		h.array[0] = h.array[len(h.array)-1]
		h.array = h.array[:len(h.array)-1]
		h.fixHeapTopDown(0)
	} else {
		node = nil
	}
	return node
}

func (h *heap) Delete(nodeTest func(Node) bool) {
	var i = 0
	for i < len(h.array) {
		if nodeTest(h.array[i]) {
			h.array[i] = h.array[len(h.array)-1]
			h.array = h.array[:len(h.array)-1]
			h.fixHeapTopDown(i)
		} else {
			i++
		}
	}
}
