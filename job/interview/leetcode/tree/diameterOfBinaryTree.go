package tree

/**
给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。

示例 :
给定二叉树

          1
         / \
        2   3
       / \
      4   5
返回3, 它的长度是路径 [4,2,1,3] 或者[5,2,1,3]。
注意：两结点之间的路径长度是以它们之间边的数目表示。
*/
func diameterOfBinaryTree(root *TreeNode) int {
	// 本质是找到某个节点的左右子树深度之和为最大
	ans := 1
	var depth func(node *TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		l, r := depth(node.Left), depth(node.Right)
		ansTemp := l + r + 1 // 1 是当前节点自身
		if ansTemp > ans {
			ans = ansTemp
		}

		if l > r {
			r = l
		}
		return r + 1
	}
	depth(root)

	return ans - 1 // 题意本质是两个节点之间的边数，而不是节点数
}
