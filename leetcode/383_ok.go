package leetcode

func CanConstruct(ransomNote string, magazine string) bool {
	ransomNoteLetters := make(map[rune]int)

	for _, letter := range ransomNote {
		ransomNoteLetters[letter]++
	}

	for _, letter := range magazine {
		ransomNoteLetters[letter]--
		if ransomNoteLetters[letter] <= 0 {
			delete(ransomNoteLetters, letter)
		}
	}

	return len(ransomNoteLetters) == 0
}
