package bst

func bstFromPreorder(preorder []int) *Node {
	var root *Node

	for _, v := range preorder {
		root = insert(root, v)
	}

	return root
}

func insert(n *Node, v int) *Node {
	if n == nil {return &Node{Val: v}}

	if n.Val > v {
		// push left
		n.Left = insert(n.Left, v)
	} else {
		n.Right = insert(n.Right, v)
	}

	return n
}