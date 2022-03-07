package tree

/** 二叉树的层平均值
给定一个非空二叉树的根节点 root , 以数组的形式返回每一层节点的平均值。与实际答案相差 10-5 以内的答案可以被接受。

示例 1：

输入：root = [3,9,20,null,null,15,7]
输出：[3.00000,14.50000,11.00000]
解释：第 0 层的平均值为 3,第 1 层的平均值为 14.5,第 2 层的平均值为 11 。
因此返回 [3, 14.5, 11] 。

示例 2:
输入：root = [3,9,20,15,7]
输出：[3.00000,14.50000,11.00000]
*/
func averageOfLevels(root *TreeNode) []float64 {
	// k为层级，v为平均值
	m := make([][]float64, 0)

	var compute func(level int, node *TreeNode)
	compute = func(level int, node *TreeNode) {
		if node == nil {
			return
		}
		if len(m) < level+1 {
			m = append(m, []float64{0, 0})
		}

		m[level][0] += float64(node.Val)
		m[level][1]++

		// 有左节点
		compute(level+1, node.Left)
		compute(level+1, node.Right)
	}

	ret := make([]float64, len(m))
	compute(0, root)
	for mi := range m {
		ret = append(ret, m[mi][0]/m[mi][1])
	}

	return ret
}
