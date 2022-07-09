package bst

import (
	"math"
)

https://leetcode.com/problems/binary-tree-maximum-path-sum/

func maxSumPath(root *Node) int {
	// initialize max as the lowest number, so the first read node will always be larger
	max := math.MinUint32
	inorder(root, &max)
	return max
}

func inorder(root *Node, max *int) int {
	if root == nil { return 0 }

	// if the left and right nodes are nil, l and r = 0
	l := inorder(root.Left, max)
	r := inorder(root.Right, max)

	// find max sum of root and its children
	// only add the child if it is > 0 
	sum := root.Val
	if l > 0 { sum += l }
	if r > 0 { sum += r }
	
	// compare the sum to the max
	*max = getMax(*max, sum)

	// if the left chain or the right chain is > 0,
	// return the root  val + left/right
	// else, just return the root val
	if getMax(l, r) > 0 {
		return getMax(l, r) + root.Val
	} else {
		return root.Val
	}
}