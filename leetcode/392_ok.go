package leetcode

// https://leetcode.com/problems/is-subsequence/?envType=study-plan-v2&envId=top-interview-150
func IsSubsequence(s string, t string) bool {
	if (len(s) == 0 && len(t) > 0) || (len(s) == 0 && len(t) == 0) {
		return true
	}

	if len(s) > 0 && len(t) == 0 {
		return false
	}

	leftPointerS := 0
	leftPointerT := 0

	for leftPointerS < len(s) && leftPointerT < len(t) {
		charS := rune(s[leftPointerS])
		charT := rune(t[leftPointerT])

		if charS == charT {
			leftPointerS++
			leftPointerT++
		} else {
			leftPointerT++
		}
	}

	return leftPointerS == len(s)
}
