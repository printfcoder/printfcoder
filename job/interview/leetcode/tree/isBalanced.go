package tree

/**
给定一个二叉树，判断它是否是高度平衡的二叉树。
本题中，一棵高度平衡二叉树定义为：
一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 。
*/
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	r, l := height(root.Right), height(root.Left)
	return abs(r-l) <= 1 && isBalanced(root.Right) && isBalanced(root.Left)
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lh, rh := height(root.Left), height(root.Right)

	if lh > rh {
		return lh + 1
	}

	return rh + 1
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
