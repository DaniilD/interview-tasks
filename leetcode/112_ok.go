package leetcode

// HasPathSum https://leetcode.com/problems/path-sum/?envType=study-plan-v2&envId=top-interview-150
func HasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	return HasPathSumDfs(root, targetSum, 0)
}

func HasPathSumDfs(root *TreeNode, targetSum int, sum int) bool {
	if root == nil {
		return false
	}

	sum += root.Val
	if root.Left == nil && root.Right == nil {
		return sum == targetSum
	}

	return HasPathSumDfs(root.Left, targetSum, sum) ||
		HasPathSumDfs(root.Right, targetSum, sum)
}
