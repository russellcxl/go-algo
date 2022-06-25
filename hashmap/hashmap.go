package main

import (
	"fmt"
)

// ----------------------------- hashing --------------------------------

/*
Overview:

	For any key (string), convert it into an index (int). Index should fall within the MAP_SIZE.
	You can do this by using modulus.

	Once you have the index, you look for the matching node in your map i.e. HashMap[index].
	At this point, given that your MAP_SIZE is small, there will be index collisions.
	For example, maybe the keys "abc" and "xyz" convert into the same index.

	To deal with this, we will loop through all the nodes of HashMap[index]
	If any of the nodes' key matches the input key, great!

*/

const MAP_SIZE = 50

type HashMap struct {
	Data []*Node
}

type Node struct {
	Key   string
	Value string
	Next  *Node
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
		h.Data[index] = &Node{Key: key, Value: value}
	} else {
		// there is a collision, start going through the linked-list
		startingNode := h.Data[index]
		for startingNode != nil {
			// the key exists, modify its value
			if startingNode.Key == key {
				startingNode.Value = value
				return
			}
			startingNode = startingNode.Next
		}
		// if the key does not exist and you've reached the last node, insert key
		startingNode = &Node{Key: key, Value: value}
	}
}

func (h *HashMap) Get(key string) (string, bool) {
	index := getIndex(key)

	// if index is empty, return false
	if h.Data[index] == nil {
		return "", false
	}

	// if index isn't empty, go through nodes and look for key
	startNode := h.Data[index]

	for startNode != nil {
		if startNode.Key == key {
			return startNode.Value, true
		}
		startNode = startNode.Next
	}

	return "", false
}

// ----------------------------- main --------------------------------

func main() {
	a := NewDict()
	a.Insert("name", "tracy")
	if value, ok := a.Get("name"); ok {
		fmt.Println(value)
	} else {
		fmt.Println("Value did not match!")
	}
}
