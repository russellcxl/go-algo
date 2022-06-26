package bst

func checkBalanced(root *Node) bool {
	if traverse(root) < 0 {return false}
	return true
}

func traverse(root *Node) int {
	// if the root is nil, return 0 to counter
	if root == nil {return 0}

	l := traverse(root.Left)
	r := traverse(root.Right)

	// if one of the subtrees below has a height difference of > 1,
	// return -1 all the way to the root i.e. the main function
	if l < 0 || r < 0 {
		return -1
	}

	// if this subtree has a height diff of > 1, return -1
	if abs(l, r) < 0 {
		return -1
	}

	// every time you move up a node, add 1 to the counter
	return max(l, r ) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a, b int) int {
	if a - b > 0 {
		return a - b
	}
	return b - a
}
