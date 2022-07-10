type Node struct{
	Val int
	Left *Node
	Right *Node
}

func getMax(a, b int) int {
	if a > b { return a }
	return b
}