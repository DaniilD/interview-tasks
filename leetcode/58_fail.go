package leetcode

// https://leetcode.com/problems/length-of-last-word/description/?envType=study-plan-v2&envId=top-interview-150
func LengthOfLastWord(s string) int {
	chars := []rune(s)
	if len(chars) <= 1 {
		return len(chars)
	}
	lastCharWordIndex := -1
	firstCharWordIndex := -1
	for i := len(chars) - 1; i != 0; i-- {
		char := chars[i]
		if string(char) != " " && lastCharWordIndex == -1 {
			lastCharWordIndex = i
		}

		if (lastCharWordIndex != -1 && firstCharWordIndex == -1 && string(char) == " ") || i == 1 {
			firstCharWordIndex = i + 1
		}
	}

	lastWord := s[firstCharWordIndex : lastCharWordIndex+1]
	chars = []rune(lastWord)

	return len(chars)
}
