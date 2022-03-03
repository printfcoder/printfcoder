package tree

/**
给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。

示例 1：

输入：root = [3,9,20,null,null,15,7]
输出：[[3],[9,20],[15,7]]
示例 2：

输入：root = [1]
输出：[[1]]
示例 3：

输入：root = []
输出：[]

提示：

树中节点数目在范围 [0, 2000] 内
-1000 <= TreeNode.val <= 1000
*/
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	dfsLevel(root, 0, &res)
	return res
}

// 在当前节点与层级、数组处理
// 处理过程中增加层级
func dfsLevel(node *TreeNode, level int, res *[][]int) {
	if node == nil {
		return
	}
	// 增加层级
	currLevel := level + 1
	// 保证数组长度足够
	for len(*res) < currLevel {
		*res = append(*res, []int{})
	}

	// 把值放进当前层所在的数组
	(*res)[currLevel-1] = append((*res)[currLevel-1], node.Val)
	dfsLevel(node.Left, currLevel, res)
	dfsLevel(node.Right, currLevel, res)
}

func bfsLevel(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var queue []*TreeNode
	queue = append(queue, root)
	var res [][]int
	var tmp []int
	curNum, nextLevelNum := 1, 0
	for len(queue) != 0 {
		if curNum > 0 {
			node := queue[0]
			if node.Left != nil {
				queue = append(queue, node.Left)
				nextLevelNum++
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
				nextLevelNum++
			}
			curNum--
			tmp = append(tmp, node.Val)
			queue = queue[1:]
		}
		if curNum == 0 {
			res = append(res, tmp)
			curNum = nextLevelNum
			nextLevelNum = 0
			tmp = []int{}
		}
	}
	return res
}
