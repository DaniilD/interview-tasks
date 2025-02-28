package leetcode

func AverageOfLevels(root *TreeNode) []float64 {
	levelMap := make(map[int][]int)
	createLevelMap(root, 0, levelMap)

	result := make([]float64, 0, len(levelMap))
	for i := 0; i < len(levelMap); i++ {
		nums := levelMap[i]
		sum := 0
		for _, num := range nums {
			sum += num
		}

		result = append(result, float64(sum)/float64(len(nums)))
	}

	return result
}

func createLevelMap(node *TreeNode, level int, levels map[int][]int) {
	if node == nil {
		return
	}

	levels[level] = append(levels[level], node.Val)
	nextLevel := level + 1

	createLevelMap(node.Left, nextLevel, levels)
	createLevelMap(node.Right, nextLevel, levels)
}
