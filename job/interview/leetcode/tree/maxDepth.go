package tree

//给定一个二叉树，找出其最大深度。
//
// 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
//
// 说明: 叶子节点是指没有子节点的节点。
//
// 示例：
//给定二叉树 [3,9,20,null,null,15,7]，
//
//     3
//   / \
//  9  20
//    /  \
//   15   7
//
// 返回它的最大深度 3 。
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 1054 👎 0

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// O(N) time, O(heightOf(N)) space
func maxDepthWithDepthFirst(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDepthWithDepthFirst(root.Left), maxDepthWithDepthFirst(root.Right)) + 1
}

// 广度优先搜索
// 使用一数组来存放当前节点的左右孩子，然后遍历这个数组，每次遍历迭代会把当前迭代的元素清掉，当当前迭代的元素有孩子时时把孩子塞到数组中，继续遍历
// 每遍历完一次说明有1级深度
// O(N) time, O(N) space
func maxDepthWithBreadthFirst(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var arr []*TreeNode
	arr = append(arr, root)
	ans := 0

	for sz := len(arr); len(arr) > 0; sz = len(arr) {
		for sz > 0 {
			// 当前节点
			cn := arr[0]

			// 当前迭代过的元素位置要去掉
			arr = arr[1:]
			sz--

			if cn.Right != nil {
				arr = append(arr, cn.Right)
			}
			if cn.Left != nil {
				arr = append(arr, cn.Left)
			}
		}

		ans++
	}

	return ans
}
