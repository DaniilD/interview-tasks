package leetcode

// MaxDepth https://leetcode.com/problems/maximum-depth-of-binary-tree/description/?envType=study-plan-v2&envId=top-interview-150
func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(MaxDepth(root.Left)+1, MaxDepth(root.Right)+1)
}

func max(one int, two int) int {
	if one > two {
		return one
	}

	return two
}
