package tree

import "math"

/** 二叉树的最小深度

给定一个二叉树，找出其最小深度。
最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
说明：叶子节点是指没有子节点的节点。

*/
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	minD := math.MaxInt
	if root.Left != nil {
		lD := minDepth(root.Left)
		if lD < minD {
			minD = lD
		}
	}

	if root.Right != nil {
		rD := minDepth(root.Right)
		if rD < minD {
			minD = rD
		}
	}

	return minD + 1
}
