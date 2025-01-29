package leetcode

// IsHappy https://leetcode.com/problems/happy-number/?envType=study-plan-v2&envId=top-interview-150
func IsHappy(n int) bool {
	existSum := make(map[int]struct{})
	sum := 0
	for sum != 1 {

		sum = 0
		for n != 0 {
			num := n % 10
			sum += num * num
			n = n / 10
		}

		if _, ok := existSum[sum]; ok {
			return false
		}
		existSum[sum] = struct{}{}

		n = sum
	}
	return true
}
