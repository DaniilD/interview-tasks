package leetcode

// TwoSum https://leetcode.com/problems/two-sum/?envType=study-plan-v2&envId=top-interview-150
func TwoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return []int{}
	}

	numsToIndex := make(map[int]int)
	for i, num := range nums {
		if indexOne, ok := numsToIndex[target-num]; ok {
			return []int{indexOne, i}
		}

		numsToIndex[num] = i
	}

	return []int{}
}
