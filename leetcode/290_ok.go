package leetcode

import (
	"strings"
)

func WordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}

	mapper := make(map[rune]string)
	useWord := make(map[string]struct{})
	for i := 0; i < len(pattern); i++ {
		charPattern := rune(pattern[i])
		word := words[i]

		if expectWord, ok := mapper[charPattern]; ok {
			if expectWord != word {
				return false
			}
		} else {
			mapper[charPattern] = word
			if _, ok := useWord[word]; ok {
				return false
			}
			useWord[word] = struct{}{}
		}
	}

	return true
}
