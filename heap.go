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
}

// Debug : Mode
var Debug bool = true

// New :
func New() Heap {
	return &heap{}
}

// FromArray :
func FromArray(nodes []Node) Heap {
	h := &heap{nodes: nodes}
	h.heapify(0)
	h.heapConsistent()
	return h
}
