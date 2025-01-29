package leetcode

func IsIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mapper := make(map[rune]rune)
	usesLetterT := make(map[rune]struct{})

	for i := 0; i < len(s); i++ {
		charS := rune(s[i])
		charT := rune(t[i])

		if charExpect, ok := mapper[charS]; ok {
			if charExpect != charT {
				return false
			}
		} else {
			mapper[charS] = charT
			if _, ok2 := usesLetterT[charT]; ok2 {
				return false
			}
			usesLetterT[charT] = struct{}{}
		}
	}

	return true
}
