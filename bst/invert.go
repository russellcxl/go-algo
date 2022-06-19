func invert(root *Node) {
	if root == nil {return}

	root.Left, root.Right = root.Right, root.left

	invert(root.Right)
	invert(root.Left)
}