package leetcode

// IsAnagram https://leetcode.com/problems/valid-anagram/description/?envType=study-plan-v2&envId=top-interview-150
func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mapper := make(map[rune]int)
	for _, char := range s {
		mapper[char]++
	}

	for _, char := range t {
		mapper[char]--
		if mapper[char] <= 0 {
			delete(mapper, char)
		}
	}

	return len(mapper) == 0
}
