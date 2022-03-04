package tree

// 二叉树的后序遍历
func postorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	var postorder func(node *TreeNode)
	postorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postorder(node.Left)
		postorder(node.Right)
		ret = append(ret, node.Val)
	}

	postorder(root)
	return ret
}
