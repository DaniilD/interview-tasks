package leetcode

// CountNodes https://leetcode.com/problems/count-complete-tree-nodes/description/?envType=study-plan-v2&envId=top-interview-150
func CountNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + CountNodes(root.Left) + CountNodes(root.Right)
}
