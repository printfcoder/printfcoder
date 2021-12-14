package main

//给定一个二叉树的根节点 root ，返回它的 中序（左顶右） 遍历。
//
// 
//
// 示例 1： 
//
// 
//输入：root = [1,null,2,3]
//输出：[1,3,2]
// 
//
// 示例 2： 
//
// 
//输入：root = []
//输出：[]
// 
//
// 示例 3： 
//
// 
//输入：root = [1]
//输出：[1]
// 
//
// 示例 4： 
//
// 
//输入：root = [1,2]
//输出：[2,1]
// 
//
// 示例 5： 
//
// 
//输入：root = [1,null,2]
//输出：[1,2]
// 
//
// 
//
// 提示： 
//
// 
// 树中节点数目在范围 [0, 100] 内 
// -100 <= Node.val <= 100

// 递归，
// O(n) time，O(n) space，空间为高度
func inorderTraversalWithRecursion(root *Node) (res []int) {
	var inorder func(node *Node)
	inorder = func(node *Node) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return
}
// 迭代
// 从根节点开始把左节点放到栈中，然后把栈中每个节点的左值和顶值放到返回值中
// 剩下的就是当前迭代的右子树，然后再把最顶右子树节点同样的方式遍历，把其顶当成根，直到所有节点都遍历到
// O(n) time，O(n) space，空间为高度
func inorderTraversalWithIteration(root *Node) (res []int) {
	var stack []*Node
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return
}
