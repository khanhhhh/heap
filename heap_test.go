package heap

import (
	"fmt"
	"math/rand"
	"testing"
)

func verify(h Heap) {
	count := 0
	last := -1
	for h.Len() > 0 {
		current := h.Pop()
		if current < last {
			panic("fail 1")
		}
		last = current
		count++
	}
}

func TestHeap(t *testing.T) {
	h := New()
	k1 := h.Push(1)
	k2 := h.Push(2)
	k3 := h.Push(3)
	k4 := h.Push(4)
	k5 := h.Push(5)
	k6 := h.Push(6)
	h.Update(k1, 5)
	h.Update(k2, 0)
	h.Update(k3, 2)
	h.Update(k4, 3)
	h.Update(k5, 3)
	h.Update(k6, 7)
	for h.Len() > 0 {
		fmt.Println(h.Pop())
	}
}

func TestHeapRandom(t *testing.T) {
	rand.Seed(1234)
	numTests := 1000
	numValues := 1000
	func() {
		for t := 0; t < numTests; t++ {
			h := New()
			for i := 0; i < numValues; i++ {
				h.Push(rand.Intn(numValues))
			}
			verify(h)
		}
	}()

	func() {
		for t := 0; t < numTests; t++ {
			values := make([]Value, numValues)
			for i := 0; i < numValues; i++ {
				values[i] = rand.Intn(numValues)
			}
			h, _, _ := FromArray(values)
			verify(h)
		}
	}()
	func() {
		for t := 0; t < numTests; t++ {
			values := make([]Value, numValues)
			for i := 0; i < numValues; i++ {
				values[i] = rand.Intn(numValues)
			}
			h, values, keys := FromArray(values)
			for i := 0; i < numValues; i++ {
				h.Update(keys[rand.Intn(numValues)], rand.Intn(numValues))
			}

			verify(h)
		}
	}()
}
