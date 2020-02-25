package heap

// Node :
type Node interface {
	Less(other Node) bool
}

// Heap :
type Heap interface {
	Len() int
	Push(node Node)
	Pop() (node Node)
	Heapify()
}

// NewHeap :
func NewHeap() (h Heap) {
	return &heap{}
}
