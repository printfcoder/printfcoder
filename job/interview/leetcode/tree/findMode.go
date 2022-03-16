package tree

/** 二叉搜索树中的众数
给你一个含重复值的二叉搜索树（BST）的根节点 root ，找出并返回 BST 中的所有 众数（即，出现频率最高的元素）。
如果树中有不止一个众数，可以按 任意顺序 返回。
假定 BST 满足如下定义：

结点左子树中所含节点的值 小于等于 当前节点的值
结点右子树中所含节点的值 大于等于 当前节点的值
左子树和右子树都是二叉搜索树

示例 1：
输入：root = [1,null,2,2]
输出：[2]

示例 2：
输入：root = [0]
输出：[0]
提示：

树中节点的数目在范围 [1, 104] 内
-105 <= Node.val <= 105
*/
/**
var base, count, maxCount int

   update := func(x int) {
       if x == base {
           count++
       } else {
           base, count = x, 1
       }
       if count == maxCount {
           answer = append(answer, base)
       } else if count > maxCount {
           maxCount = count
           answer = []int{base}
       }
   }

   var dfs func(*TreeNode)
   dfs = func(node *TreeNode) {
       if node == nil {
           return
       }
       dfs(node.Left)
       update(node.Val)
       dfs(node.Right)
   }
   dfs(root)
   return
*/
func findMode(root *TreeNode) []int {
	ret, lastNum, count, maxCount := make([]int, 0), 0, 0, 0
	var max func(node *TreeNode)
	max = func(node *TreeNode) {
		if node == nil {
			return
		}

		// 使用中序
		max(node.Left)

		// 如果值和最上一个相同，则当前总数加1
		if node.Val == lastNum {
			count++
		} else {
			// 如果值和最上一个不相同，重置
			count = 1
			lastNum = node.Val
		}

		if count > maxCount {
			maxCount = count
			ret = []int{lastNum}
		} else if count == maxCount {
			ret = append(ret, lastNum)
		}

		max(node.Right)
	}

	max(root)
	return ret
}
