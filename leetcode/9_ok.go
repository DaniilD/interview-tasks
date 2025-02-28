package leetcode

func IsPalindromeNum(x int) bool {
	if x < 0 {
		return false
	}

	t := x
	rev := 0
	for t != 0 {
		mod := t % 10
		rev = rev*10 + mod

		t = t / 10
	}

	return rev == x
}
