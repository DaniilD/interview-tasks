package leetcode

// RemoveDuplicate
// https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/?envType=study-plan-v2&envId=top-interview-150
func RemoveDuplicate(n []int) int {
	count := 0
	i := 0
	for i < len(n) {
		v := n[i]
		has := false
		for j := 0; j < count; j++ {
			v2 := n[j]
			if v == v2 {
				has = true
			}
		}

		if !has {
			n[count] = v
			count++
		}

		i++
	}
	return count
}
