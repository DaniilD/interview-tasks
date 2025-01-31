package leetcode

func IsValid(s string) bool {

	brackets := make(map[rune]rune, 3)
	brackets['('] = ')'
	brackets['{'] = '}'
	brackets['['] = ']'

	stack := make([]rune, 0)
	for _, bracket := range s {
		if closeBracket, ok := brackets[bracket]; ok {
			stack = append(stack, closeBracket)
		} else {
			if len(stack) == 0 {
				return false
			}
			closeBracket := stack[len(stack)-1]
			if bracket != closeBracket {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) > 0 {
		return false
	}

	return true
}
