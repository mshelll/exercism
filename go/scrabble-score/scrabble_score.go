package scrabble

import (
	"unicode"
)

type ScrabbleMap map[rune]int

var letter_values ScrabbleMap

func init() {

	letter_values = ScrabbleMap{}

	for _, key := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {

		switch key {

			case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			   letter_values[key] = 1
			case 'D', 'G':
				letter_values[key] = 2
			case 'B', 'C', 'M', 'P':
				letter_values[key] = 3
			case 'F', 'H', 'V', 'W', 'Y':
				letter_values[key] = 4
			case 'K':
				letter_values[key] = 5
			case 'J', 'X':
				letter_values[key] = 8
			case 'Q', 'Z':
				letter_values[key] = 10
		}
	}

}

func Score(in string) (n int) {

	score := 0

	for _, value := range in {

	   score += letter_values[unicode.ToUpper(value)]
	}

	return score
}

