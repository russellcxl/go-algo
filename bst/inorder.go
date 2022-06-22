package bst

type Node struct{
	Val int
	Left *Node
	Right *Node
}

func inorder(root *Node) []int {
	
	// inorder: left > root > right > ...

	/*
					A
			B				C
		D		E	  	F		G

		always return the triangle to the previous call
		e.g. [F, C, G] will be returned to inorder(A.Right)
	*/

	if root == nil {
		return []int{}	
	}

	l := inorder(root.Left)
	r := inorder(root.Right)

	// append root, then right
	l = append(l, root.Val)
	l = append(l, r...)

	return l
	
}


