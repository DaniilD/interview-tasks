package leetcode

// IsSameTree https://leetcode.com/problems/same-tree/?envType=study-plan-v2&envId=top-interview-150
func IsSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}

	if p.Val != q.Val {
		return false
	}

	return IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
}
