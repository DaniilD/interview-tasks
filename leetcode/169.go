package leetcode

import (
	"math"
)

// [2,2,1,1,1,2,2]
// https://leetcode.com/problems/majority-element/?envType=study-plan-v2&envId=top-interview-150
func MajorityElement(nums []int) int {
	m := make(map[int]int, len(nums))
	count := math.Round(float64(len(nums)) / 2.0)

	for _, n := range nums {
		m[n]++
		if m[n] >= int(count) {
			return n
		}
	}

	return 0
}
