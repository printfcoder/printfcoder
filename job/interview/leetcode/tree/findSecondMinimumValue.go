package tree

/** 二叉树中第二小的节点
给定一个非空特殊的二叉树，每个节点都是正数，并且每个节点的子节点数量只能为 2 或 0。如果一个节点有两个子节点的话，那么该节点的值等于两个子节点中较小的一个。
更正式地说，即 root.val = min(root.left.val, root.right.val) 总成立。
给出这样的一个二叉树，你需要输出所有节点中的 第二小的值 。
如果第二小的值不存在的话，输出 -1 。

示例 1：
输入：root = [2,2,5,null,null,5,7]
输出：5
解释：最小的值是 2 ，第二小的值是 5 。

示例 2：
输入：root = [2,2,2]
输出：-1
解释：最小的值是 2, 但是不存在第二小的值。

提示：

树中节点数目在范围 [1, 25] 内
1 <= Node.val <= 231 - 1
对于树中每个节点 root.val == min(root.left.val, root.right.val)
*/
func findSecondMinimumValue(root *TreeNode) int {
	// 按题意：任一根节点的值一定是它两个子节点中最小的值
	// 所以，根节点就是整棵树最小的值，
	// 记录两个值，根值和第二小值，根值为最小，小于根值的第一个，即为我们需要的值
	// 第二小值默认为-1，这是满足题意有找不到的情况
	rootVal, ret := root.Val, -1

	var findTheNum func(node *TreeNode)
	findTheNum = func(node *TreeNode) {
		if node == nil || (ret != -1 && node.Val >= ret) {
			return
		}

		if node.Val > rootVal {
			ret = node.Val
		}

		findTheNum(node.Left)
		findTheNum(node.Right)
	}

	findTheNum(root)
	return ret
}
