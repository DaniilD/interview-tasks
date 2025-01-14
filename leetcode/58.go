package leetcode

import "fmt"

func LengthOfLastWord(s string) int {
	chars := []rune(s)
	lastCharWordIndex := -1
	firstCharWordIndex := -1
	for i := len(chars) - 1; i != 0; i-- {
		char := chars[i]
		if string(char) != " " && lastCharWordIndex == -1 {
			lastCharWordIndex = i
		}

		if lastCharWordIndex != -1 && firstCharWordIndex == -1 && string(char) == " " {
			firstCharWordIndex = i + 1
		}
	}

	fmt.Println(firstCharWordIndex)
	fmt.Println(lastCharWordIndex)
	lastWord := s[firstCharWordIndex:lastCharWordIndex]
	chars = []rune(lastWord)

	return len(chars)
}
