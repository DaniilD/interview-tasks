package leetcode

import "fmt"

// SummaryRanges https://leetcode.com/problems/summary-ranges/description/?envType=study-plan-v2&envId=top-interview-150
func SummaryRanges(nums []int) []string {
	leftPointer := 0
	rightPointer := 0

	results := make([]string, 0)
	for rightPointer < len(nums) {
		leftNum := nums[leftPointer]
		rightNum := nums[rightPointer]

		if ((rightPointer + 1) < len(nums)) && (rightNum+1 == nums[rightPointer+1]) {
			rightPointer++
		} else {
			var result string
			if leftPointer == rightPointer {
				result = fmt.Sprintf("%d", leftNum)
			} else {
				result = fmt.Sprintf("%d->%d", leftNum, rightNum)

			}

			results = append(results, result)
			leftPointer = rightPointer + 1
			rightPointer = leftPointer
		}
	}

	return results
}
