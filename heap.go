package heap

// Value :
type Value = int

// Key :
type Key = *int

// Debug : Mode
const Debug bool = true

// Heap :
type Heap interface {
	Len() int
	Push(value Value) (key Key)
	Pop() (value Value)
	Update(key Key, value Value)
}

// New :
func New() Heap {
	return &heap{}
}

// FromArray :
func FromArray(array []Value) (Heap, []Value, []Key) {
	h := &heap{
		value: make([]Value, len(array)),
		key:   make([]Key, len(array)),
	}
	defer h.heapConsistentAssert()
	for i := 0; i < len(array); i++ {
		h.value[i] = array[i]
		index := i
		h.key[i] = &index
	}
	h.heapify(0)
	return h, h.value, h.key
}
