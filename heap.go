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
	Delete(nodeTest func(Node) bool)
}

// NewHeap :
func NewHeap() (h Heap) {
	return &heap{}
}
