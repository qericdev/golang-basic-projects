package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Insert tha phrase or word to verify if it is a palindrome:")
	reader := bufio.NewReader(os.Stdin)
	phrase, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading the phrase", err)
		return
	}

	phrase = strings.TrimSpace(phrase)
	fmt.Println(isPalindrome(phrase))

}

func isPalindrome(phrase string) bool {
	// Lower letters of the phrase
	phrase = strings.ToLower(phrase)
	// Remove spaces
	phrase = strings.Join(strings.Fields(phrase), "")
	// Removing all characters that are not letters
	finalPhrase := ""

	for i := 0; i < len(phrase); i++ {
		if phrase[i] >= 'a' && phrase[i] <= 'z' {
			finalPhrase += string(phrase[i])
		}
	}

	for i := 0; i < len(finalPhrase)/2; i++ {
		if phrase[i] != phrase[len(finalPhrase)-1-i] {
			return false
		}
	}
	return true
}
