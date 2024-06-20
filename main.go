package main

import (
	"fmt"
)

func main() {
	text := "This is a test string!"
	phoneticString := phonetic.ToPhonetic(text)
	fmt.Println(phoneticString) // Output: TANGO HOTEL INDIA SIERRA INDIA SIERRA ALPHA TANGO ECHO SIERRA TANGO SIERRA ROMEO INDIA NOVEMBER GOLF !
}
