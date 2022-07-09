package bst

func hasPathSum(root *Node, targetSum int) bool {
	// if you've gone past the last leaf, return false up the callback chain
    if root == nil { return false }

	// every time you traverse a node, reduce the target value
	remainder := targetSum - root.Val

	// if you're at the leaf node and remainder == 0, return true up the callback chain
	if root.Left == nil && root.Right == nil && remainder == 0 { return true }

	// if you're not at the leaf node, call the function again
	return hasPathSum(root.Left, remainder) || hasPathSum(root.Right, remainder)
    
}