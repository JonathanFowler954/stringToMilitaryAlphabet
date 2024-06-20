package MilitaryPhonetic

import (
	"strings"
)

// Define the phonetic alphabet map
var phoneticAlphabet = map[rune]string{
	'A': "Alpha", 'B': "Bravo", 'C': "Charlie", 'D': "Delta",
	'E': "Echo", 'F': "Foxtrot", 'G': "Golf", 'H': "Hotel",
	'I': "India", 'J': "Juliett", 'K': "Kilo", 'L': "Lima",
	'M': "Mike", 'N': "November", 'O': "Oscar", 'P': "Papa",
	'Q': "Quebec", 'R': "Romeo", 'S': "Sierra", 'T': "Tango",
	'U': "Uniform", 'V': "Victor", 'W': "Whiskey", 'X': "X-ray",
	'Y': "Yankee", 'Z': "Zulu",
	'0': "Zero", '1': "One", '2': "Two", '3': "Three",
	'4': "Four", '5': "Five", '6': "Six", '7': "Seven",
	'8': "Eight", '9': "Nine", ' ': "Space",
}

// ConvertTextToPhonetic converts a string to its phonetic alphabet equivalent
func ConvertTextToPhonetic(input string) string {
	var result []string
	input = strings.ToUpper(input)
	for _, char := range input {
		if val, ok := phoneticAlphabet[char]; ok {
			result = append(result, val)
		} else {
			result = append(result, string(char))
		}
	}
	return strings.Join(result, " ")
}
