package bst

func buildTree(preorder, inorder []int) *Node {
	if len(preorder) == 0 {return nil}

	// start with the root in preorder
	root := &Node{Val: preorder[0]}

	// loop through inorder to find the root node in preorder
	// elements on the left belong to the left tree
	// elements on the right belong to the right tree
	for i := 0 ; i < len(inorder) ; i++ {
		if preorder[0] == inorder[i] {
			root.Left = buildTree(preorder[1:i+1], inorder[:i])
			root.Right = buildTree(preorder[i+1:], inorder[i+1:])
			break
		}
	}

	return root

}