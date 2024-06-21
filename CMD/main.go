package main

import (
	"bufio"
	"fmt"
	MP "github.com/jonathanfowler954/stringToMilitaryAlphabet"
	"os"
	//"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string to convert to the phonetic alphabet: ")
	input, _ := reader.ReadString('\n')
	//input = strings.TrimSpace(input)

	phonetic := MP.ConvertTextToPhonetic(input)
	fmt.Println("Phonetic Alphabet:", phonetic)
}
