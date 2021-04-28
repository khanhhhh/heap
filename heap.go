package heap

// Value :
type Value = int

// Key :
type Key *int

// Debug : Mode
const Debug bool = true

// Heap :
type Heap interface {
	Len() int
	Push(value Value, data interface{}) (key Key)
	Pop() (value Value, data interface{})
	Update(key Key, value Value)
}

// New :
func New() Heap {
	return &heap{}
}

// FromArray :
func FromArray(array []Value) Heap {
	h := &heap{
		node: make([]node, 0, len(array)),
	}
	defer h.heapConsistentAssert()
	for i := range array {
		h.node = append(h.node, newNode(len(h.node), array[i], nil))
	}
	h.heapify(0)
	return h
}
