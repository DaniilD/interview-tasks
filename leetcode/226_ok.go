package leetcode

// InvertTree https://leetcode.com/problems/invert-binary-tree/?envType=study-plan-v2&envId=top-interview-150
func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	tmp := root.Right
	root.Right = InvertTree(root.Left)
	root.Left = InvertTree(tmp)

	return root
}
