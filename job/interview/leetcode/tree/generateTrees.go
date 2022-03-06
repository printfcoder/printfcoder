package tree

/** 不同的二叉搜索树 II
给你一个整数 n ，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树 。可以按 任意顺序 返回答案。

示例 1：
输入：n = 3
输出：[[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]
示例 2：

输入：n = 1
输出：[[1]]
*/
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	// 两边递进，逐个生成
	return setLeftAndRight(1, n)
}

func setLeftAndRight(lower, greater int) []*TreeNode {
	if lower > greater {
		return []*TreeNode{nil}
	}
	var allTrees []*TreeNode
	// 枚举可行根节点
	for i := lower; i <= greater; i++ {
		// 左子树
		leftTrees := setLeftAndRight(lower, i-1)
		// 右子树
		rightTrees := setLeftAndRight(i+1, greater)
		// 拼接
		for _, left := range leftTrees {
			for _, right := range rightTrees {
				currTree := &TreeNode{i, nil, nil}
				currTree.Left = left
				currTree.Right = right
				allTrees = append(allTrees, currTree)
			}
		}
	}
	return allTrees
}
