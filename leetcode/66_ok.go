package leetcode

// PlusOne https://leetcode.com/problems/plus-one/?envType=study-plan-v2&envId=top-interview-150
func PlusOne(digits []int) []int {
	adder := 1
	for i := len(digits) - 1; i >= 0; i-- {
		num := digits[i]
		num = num + adder
		if num == 10 {
			num = 0
			adder = 1
		} else {
			adder = 0
		}

		digits[i] = num
	}

	if adder > 0 {
		return append([]int{adder}, digits...)
	}

	return digits
}
