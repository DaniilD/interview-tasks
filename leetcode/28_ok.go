package leetcode

import "fmt"

// https://leetcode.com/problems/longest-common-prefix/description/?envType=study-plan-v2&envId=top-interview-150
func StrStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}

	if len(haystack) == 0 {
		return -1
	}

	for i, char := range haystack {
		if char == rune(needle[0]) {
			if i+len(needle) > len(haystack) {
				return -1
			}
			subStr := haystack[i : i+len(needle)]
			fmt.Println(subStr)
			if subStr == needle {
				return i
			}
		}
	}

	return -1
}
