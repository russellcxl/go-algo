package main

import (
	"fmt"
)

// ----------------------------- parameters and models --------------------------------

const MAP_SIZE = 50

type HashMap struct {
	Data []*Node
}

type Node struct {
	key   string
	value string
	next  *Node
}

func NewDict() *HashMap {
	return &HashMap{
		Data: make([]*Node, MAP_SIZE),
	}
}

func hash(key string) (hash uint32) {
	hash = 0
	for _, ch := range key {
		hash += uint32(ch)
		hash += hash << 10
		hash ^= hash >> 6
	}

	hash += hash << 3
	hash ^= hash >> 11
	hash += hash << 15

	return
}

func getIndex(key string) (index int) {
	return int(hash(key)) % MAP_SIZE
}

func (h *HashMap) Insert(key string, value string) {
	index := getIndex(key)

	if h.Data[index] == nil {
		// index is empty, go ahead and insert
		h.Data[index] = &Node{key: key, value: value}
	} else {
		// there is a collision, get into linked-list mode
		starting_node := h.Data[index]
		for ; starting_node.next != nil; starting_node = starting_node.next {
			if starting_node.key == key {
				// the key exists, its a modifying operation
				starting_node.value = value
				return
			}
		}
		starting_node.next = &Node{key: key, value: value}
	}
}

// ----------------------------- main --------------------------------

func main() {
	//a := NewDict()
	//fmt.Printf("%v+", a)

	fmt.Printf("%08b\n%08b", hash("a"), hash("b"))
}
