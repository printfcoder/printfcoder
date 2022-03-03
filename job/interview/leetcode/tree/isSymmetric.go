package tree

// O(n) time，O(n) space，空间为高度
func isSymmetricWithRecursion(root *TreeNode) bool {
	return check(root, root)
}

// O(n) time，O(n) space，空间为高度
func isSymmetricWithIteration(root *TreeNode) bool {
	var q []*TreeNode
	q = append(q, root)
	q = append(q, root)
	u, v := root, root
	for len(q) > 0 {
		u, v = q[0], q[1]
		q = q[2:]
		if u == nil && v == nil {
			continue
		}
		if u == nil || v == nil {
			return false
		}

		if u.Val != v.Val {
			return false
		}

		q = append(q, u.Left)
		q = append(q, v.Right)

		q = append(q, u.Right)
		q = append(q, v.Left)
	}

	return true
}

func check(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
}
