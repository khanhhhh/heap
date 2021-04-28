package heap

import "fmt"

type node struct {
	key   Key
	value Value
	data  interface{}
}

func newNode(index int, value Value, data interface{}) node {
	keyHolder := index
	return node{
		key:   &keyHolder,
		value: value,
		data:  data,
	}
}

type heap struct {
	node []node
}

func (h *heap) heapConsistentAssert() {
	if Debug && len(h.node) > 1 {
		for parent := 0; parent < len(h.node); parent++ {
			left, right := childrenOf(parent)
			if (h.isInHeap(left) && h.node[left].value < h.node[parent].value) || (h.isInHeap(right) && h.node[right].value < h.node[parent].value) {
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
	return left >= len(h.node)
}

func (h *heap) isInHeap(index int) bool {
	return index >= 0 && index < len(h.node)
}

func (h *heap) fixHeapTopDown(index int) {
	if h.isLeafNode(index) {
		return
	}
	left, right := childrenOf(index)
	var less = left
	if h.isInHeap(right) && h.node[right].value < h.node[left].value {
		less = right
	}
	if h.node[less].value >= h.node[index].value {
		return
	}

	h.node[index], h.node[less] = h.node[less], h.node[index]
	*h.node[index].key = index
	*h.node[less].key = less
	h.fixHeapTopDown(less)
}

func (h *heap) fixHeapBottomUp(index int) {
	if index <= 0 {
		return
	}
	parent := parentOf(index)
	left, right := childrenOf(parent)
	var less = left
	if h.isInHeap(right) && h.node[right].value < h.node[left].value {
		less = right
	}
	if h.node[less].value >= h.node[parent].value {
		return
	}
	h.node[parent], h.node[less] = h.node[less], h.node[parent]
	*h.node[parent].key = parent
	*h.node[less].key = less
	h.fixHeapBottomUp(parent)
}

func (h *heap) Len() int {
	return len(h.node)
}

func (h *heap) Push(value Value, data interface{}) (key Key) {
	defer h.heapConsistentAssert()
	h.node = append(h.node, newNode(len(h.node), value, data))
	key = h.node[len(h.node)-1].key
	h.fixHeapBottomUp(len(h.node) - 1)
	return key
}

func (h *heap) Pop() (value Value, data interface{}) {
	defer h.heapConsistentAssert()
	if len(h.node) <= 0 {
		panic("heap empty")
	}
	node := h.node[0]
	h.node[0], h.node = h.node[len(h.node)-1], h.node[:len(h.node)-1]
	if len(h.node) > 0 {
		*h.node[0].key = 0
	}
	h.fixHeapTopDown(0)
	return node.value, node.data
}

func (h *heap) Update(key Key, value Value) {
	defer h.heapConsistentAssert()
	old := h.node[*key].value
	h.node[*key].value = value
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
