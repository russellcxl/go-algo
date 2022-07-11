package bst

func postorder(root *Node) []int {

	// left, right, root
	// almost like reverse bfs

	// create 2 stacks -- `stack` and `out`
	// return the inversion of stack 2
	// push root into stack 1
	// pop root into stack 2
	// push root.Left, root.Right into stack 1
	// pop from 1, push into 2
	// ...
	// you'll eventually end up with the second stack as an inverted list of postorder
	// e.g. if postorder = [1, 2, 3, 4, 5], second stack will be [5, 4, 3, 2, 1]

	if root == nil {
		return []int{}
	}

	stack := []*Node{root}
	var out []int

	for len(stack) != 0 {
		popped := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]

		out = append(out, popped.Val)

		if popped.Left != nil {
			stack = append(stack, popped.Left)
		}

		if popped.Right != nil {
			stack = append(stack, popped.Right)
		}
	}

	var ans []int
	for i := len(out) - 1 ; i >= 0 ; i-- {
		ans = append(ans, out[i])
	}
	return ans
}