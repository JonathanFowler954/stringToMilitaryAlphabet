package MilitaryPhonetic

import (
	"strings"
)

// MilitaryPhoneticAlphabet is a map containing the letter-to-phonetic equivalents
var Alphabet = map[rune]string{
	'A': "Alpha", 'B': "Bravo", 'C': "Charlie", 'D': "Delta", 'E': "Echo",
	'F': "Foxtrot", 'G': "Golf", 'H': "Hotel", 'I': "India", 'J': "Juliet",
	'K': "Kilo", 'L': "Lima", 'M': "Mike", 'N': "November", 'O': "Oscar",
	'P': "Papa", 'Q': "Quebec", 'R': "Romeo", 'S': "Sierra", 'T': "Tango",
	'U': "Uniform", 'V': "Victor", 'W': "Whiskey", 'X': "X-ray", 'Y': "Yankee",
	'Z': "Zulu",
}

// ToMilitaryPhonetic converts a string to its military phonetic alphabet representation
func ToMilitaryPhonetic(text string) string {
	phoneticText := []string{}
	for _, char := range strings.ToUpper(text) {
		if phonetic, ok := Alphabet[char]; ok {
			phoneticText = append(phoneticText, phonetic)
		} else {
			phoneticText = append(phoneticText, string(char)) // Handle non-alphabetic characters
		}
	}
	return strings.Join(phoneticText, " ")
}
