package leetcode

// IsSymmetric https://leetcode.com/problems/symmetric-tree/?envType=study-plan-v2&envId=top-interview-150
func IsSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}

	return IsSymmetricDfs(root.Left, root.Right)
}

func IsSymmetricDfs(one *TreeNode, two *TreeNode) bool {
	if one == nil || two == nil {
		return one == two
	}

	if one.Val != two.Val {
		return false
	}

	return IsSymmetricDfs(one.Right, two.Left) && IsSymmetricDfs(one.Left, two.Right)
}
