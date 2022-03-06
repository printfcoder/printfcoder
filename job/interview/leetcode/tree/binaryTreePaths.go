package tree

import "strconv"

/** 二叉树的所有路径
给你一个二叉树的根节点 root ，按 任意顺序 ，返回所有从根节点到叶子节点的路径。

叶子节点 是指没有子节点的节点。

示例 1：

输入：root = [1,2,3,null,5]
输出：["1->2->5","1->3"]
示例 2：

输入：root = [1]
输出：["1"]

提示：

树中节点的数目在范围 [1, 100] 内
-100 <= Node.val <= 100
*/
func binaryTreePaths(root *TreeNode) []string {
	paths := make([]string, 0)

	var path func(pathStr string, node *TreeNode)
	path = func(pathStr string, node *TreeNode) {
		if node == nil {
			return
		}
		pathStr += strconv.Itoa(node.Val)
		if node.Left == nil && node.Right == nil {
			paths = append(paths, pathStr)
		} else {
			pathStr += "->"
			path(pathStr, node.Right)
			path(pathStr, node.Left)
		}
	}

	path("", root)
	return paths
}
