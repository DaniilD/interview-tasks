package leetcode

import "strings"

// https://leetcode.com/problems/valid-palindrome/?envType=study-plan-v2&envId=top-interview-150
func IsPalindrome(s string) bool {
	if len(s) == 0 || s == " " {
		return true
	}

	leftPointer := 0
	rightPointer := len(s) - 1

	allowSymbols := map[rune]struct{}{
		'q': {}, 'w': {}, 'e': {}, 'r': {}, 't': {}, 'y': {}, 'u': {}, 'i': {}, 'o': {},
		'p': {}, 'a': {}, 's': {}, 'd': {}, 'f': {}, 'g': {}, 'h': {}, 'j': {}, 'k': {},
		'l': {}, 'z': {}, 'x': {}, 'c': {}, 'v': {}, 'b': {}, 'n': {}, 'm': {}, '0': {},
		'1': {}, '2': {}, '3': {}, '4': {}, '5': {}, '6': {}, '7': {}, '8': {}, '9': {},
	}

	s = strings.ToLower(s)
	isPalindrome := true
	for leftPointer <= rightPointer {
		leftChar := rune(s[leftPointer])
		rightChar := rune(s[rightPointer])

		if _, ok := allowSymbols[leftChar]; !ok {
			leftPointer++
			continue
		}

		if _, ok := allowSymbols[rightChar]; !ok {
			rightPointer--
			continue
		}

		if leftChar != rightChar {
			isPalindrome = false
		}

		leftPointer++
		rightPointer--
	}

	return isPalindrome
}
