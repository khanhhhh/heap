package heap

import "fmt"

type heap struct {
	value []Value
	key   []*int
}

func (h *heap) heapConsistentAssert() {
	if Debug && len(h.value) > 1 {
		for parent := 0; parent < len(h.value); parent++ {
			left, right := childrenOf(parent)
			if (h.isInHeap(left) && h.value[left] < h.value[parent]) || (h.isInHeap(right) && h.value[right] < h.value[parent]) {
				fmt.Printf("parent %d left %d right %d\n", parent, left, right)
				panic("heap inconsistent")
			}
		}
	}
}

func childrenOf(parent int) (left int, right int) {
	return (2 * (parent + 1)) - 1, 2 * (parent + 1)
}

func parentOf(child int) (parent int) {
	return (child+1)/2 - 1
}

func (h *heap) isLeafNode(index int) bool {
	left, _ := childrenOf(index)
	return left >= len(h.value)
}

func (h *heap) isInHeap(index int) bool {
	return index >= 0 && index < len(h.value)
}

func (h *heap) fixHeapTopDown(index int) {
	if h.isLeafNode(index) {
		return
	}
	left, right := childrenOf(index)
	var less int = left
	if h.isInHeap(right) && h.value[right] < h.value[left] {
		less = right
	}
	if h.value[less] >= h.value[index] {
		return
	}
	h.value[index], h.value[less] = h.value[less], h.value[index]
	h.key[index], h.key[less] = h.key[less], h.key[index]
	*h.key[index] = index
	*h.key[less] = less
	h.fixHeapTopDown(less)
}

func (h *heap) fixHeapBottomUp(index int) {
	if index <= 0 {
		return
	}
	parent := parentOf(index)
	left, right := childrenOf(parent)
	var less int = left
	if h.isInHeap(right) && h.value[right] < h.value[left] {
		less = right
	}
	if h.value[less] >= h.value[parent] {
		return
	}
	h.value[parent], h.value[less] = h.value[less], h.value[parent]
	h.key[parent], h.key[less] = h.key[less], h.key[parent]
	*h.key[parent] = parent
	*h.key[less] = less
	h.fixHeapBottomUp(parent)
}

func (h *heap) Len() int {
	return len(h.value)
}

func (h *heap) Push(value Value) (key Key) {
	defer h.heapConsistentAssert()
	h.value = append(h.value, value)
	index := len(h.value) - 1
	h.key = append(h.key, &index)
	h.fixHeapBottomUp(index)
	return &index
}

func (h *heap) Pop() (value Value) {
	defer h.heapConsistentAssert()
	if len(h.value) <= 0 {
		panic("heap empty")
	}
	value = h.value[0]
	h.value[0], h.value = h.value[len(h.value)-1], h.value[:len(h.value)-1]
	h.key[0], h.key = h.key[len(h.key)-1], h.key[:len(h.key)-1]
	if len(h.key) > 0 {
		*h.key[0] = 0
	}
	h.fixHeapTopDown(0)
	return value
}

func (h *heap) Update(key Key, value Value) {
	defer h.heapConsistentAssert()
	old := h.value[*key]
	h.value[*key] = value
	switch {
	case value < old:
		h.fixHeapBottomUp(*key)
	case value > old:
		h.fixHeapTopDown(*key)
	}
}

func (h *heap) heapify(index int) {
	if h.isLeafNode(index) {
		return
	}
	left, right := childrenOf(index)
	h.heapify(left)
	h.heapify(right)
	h.fixHeapTopDown(index)
}
