package tree

/** 二叉搜索树的范围和
给定二叉搜索树的根结点 root，返回值位于范围 [low, high] 之间的所有结点的值的和。

示例 1：
输入：root = [10,5,15,3,7,null,18], low = 7, high = 15
输出：32

示例 2：
输入：root = [10,5,15,3,7,13,18,1,null,6], low = 6, high = 10
输出：23

提示：
树中节点数目在范围 [1, 2 * 104] 内
1 <= Node.val <= 105
1 <= low <= high <= 105
所有 Node.val 互不相同
*/
func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	if root.Val > high {
		return rangeSumBST(root.Left, low, high)
	}
	if root.Val < low {
		return rangeSumBST(root.Right, low, high)
	}

	return root.Val + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
}

func rangeSumBST2(root *TreeNode, low int, high int) int {
	queue := []*TreeNode{root}
	ans := 0
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		if n == nil {
			continue
		}
		// 大于，则找左边的进入队列
		if n.Val > high {
			queue = append(queue, n.Left)
		} else if n.Val < low { // 小于，则找右边的进入队列
			queue = append(queue, n.Right)
		} else {
			ans += n.Val
			queue = append(queue, n.Left, n.Right)
		}
	}

	return ans
}
