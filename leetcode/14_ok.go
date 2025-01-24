package leetcode

import (
	"sort"
)

// https://leetcode.com/problems/longest-common-prefix/description/?envType=study-plan-v2&envId=top-interview-150
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})

	i := 0
	result := ""
	if len(strs[0]) == 0 {
		return result
	}

	for i < len(strs[0]) {
		char := rune(strs[0][i])
		isCharEqual := true
		for _, str := range strs[1:] {
			if char != rune(str[i]) {
				isCharEqual = false
				break
			}
		}

		if isCharEqual {
			result += string(char)
		} else {
			break
		}
		i++
	}

	return result
}
