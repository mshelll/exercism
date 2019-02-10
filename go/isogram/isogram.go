package isogram

import (
	"unicode"
)

type letterMap map[rune]int

var letters letterMap

func IsIsogram(in string) (out bool) {

	letters = letterMap{}

	for _, char := range in {

		index := unicode.ToLower(char)

		if char == ' ' || char == '-' {
			continue
		} else if letters[index] == 0 {
			letters[index] = 1
		} else {
			return false
		}
	}
	return true
}
