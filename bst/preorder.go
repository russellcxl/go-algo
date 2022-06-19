package bst

/*
type Node struct{
	Val int
	Left *Node
	Right *Node
}
*/

func preorder(root *Node) []int {
	// iterative preorder traversal: root, left, right

	// create a stack
	// push root into stack
	// for stack != empty
	// pop node > print
	// push node.Right, push node.Left
	// pop node > print
	// ...

	var out []int
	stack := []*Node{root}

	for len(stack) != 0 {
		popped := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]

		out = append(out, popped.Val)

		if popped.Right != nil {
			stack = append(stack, popped.Right)
		}

		if popped.Left != nil {
			stack = append(stack, popped.Left)
		}
	}

	return out
}