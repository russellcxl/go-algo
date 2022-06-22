package bst

func invert(root *Node) {
	if root == nil {return}

	root.Left, root.Right = root.Right, root.Left

	invert(root.Right)
	invert(root.Left)
}