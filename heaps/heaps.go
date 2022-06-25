package heaps

import "errors"

type Heapify interface {
	Insert()
	Extract()
}

type MaxHeap struct {
	array []int
}

// Insert inserts a key into the heap
func (m *MaxHeap) Insert(key int) {
	m.array = append(m.array, key)
	m.maxHeapifyUp(len(m.array) - 1)
}

// Extract returns the largest key i.e. the root
func (m *MaxHeap) Extract() (int, error) {
	if len(m.array) == 0 {
		return 0, errors.New("no elements in the heap")
	}
	last := len(m.array) - 1
	extracted := m.array[0]

	// set the last as the root, and shorten the array
	m.array[0] = m.array[last]
	m.array = m.array[:last]

	m.maxHeapifyDown(0)

	return extracted, nil
}

func (m *MaxHeap) maxHeapifyUp(i int) {
	// while parent is smaller than current
	// swap parent and current
	// move pointer to parent
	for m.array[parent(i)] < m.array[i] {
		m.swap(parent(i), i)
		i = parent(i)
	}
}

func (m *MaxHeap) maxHeapifyDown(i int) {
	lastIndex := len(m.array) - 1

	l, r := left(i), right(i)

	for l <= lastIndex {
		var childToCompare int

		// if left child > right
		if l == lastIndex || m.array[l] > m.array[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		// if current val < childToCompare, swap
		// move pointer to child and its children nodes
		if m.array[i] < m.array[childToCompare] {
			m.swap(i, childToCompare)
			i, l, r = childToCompare, left(i), right(i)
		} else {
			return
		}
	}

}

// e.g. if root = i, left = 2i + 1, right = 2i + 2
// gets parent idx from left child idx
func parent(i int) int {
	return (i - 1) / 2
}

// gets left idx from parent idx
func left(i int) int {
	return (i * 2) + 1
}

// gets right idx from parent idx
func right(i int) int {
	return (i * 2) + 2
}

func (m *MaxHeap) swap(a, b int) {
	m.array[a], m.array[b] = m.array[b], m.array[a]
}
