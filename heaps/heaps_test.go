package heaps

import (
	"fmt"
	"testing"
)

func TestHeaps(t *testing.T) {
	m := &MaxHeap{}
	buildHeap := []int{10, 20, 30, 5, 45, 87, 78, 12, 55}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Printf("inserted %d; new: %v\n", v, m.array)
	}

	for i := 0; i < 5; i++ {
		extracted, err := m.Extract()
		if err != nil {
			t.Error(err)
		}
		fmt.Printf("extracted: %d; new: %v\n", extracted, m.array)
	}
}
