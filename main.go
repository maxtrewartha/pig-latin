package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	// Gets the array from position 1 (position 0 is the name of the application)
	phrase := os.Args[1:]
	fmt.Println(pigLatin(phrase))
}

// Yes, this is a function, there really isnt anything to say about it
func isVowel(letter string) bool {
	r, _ := regexp.Compile("^[aeiou]$")
	return r.MatchString(letter)
}

func pigLatin(phrase []string) string {
	// It makes
	for i := range phrase {
		/**
		 * This takes the first consonant cluster from the start of a word and then moves it to the end of the word and then adds "way"/"ay"
		 * Eg. food -> oodfay, swimming -> immingsway
		 */

		// If it starts with a vowel, just add "way"
		if isVowel(string(phrase[i][0])) {
			phrase[i] += "way"
			continue
		}
		// If its a small word and doesnt start with a vowel, just add "ay"
		if len(phrase[i]) < 2 {
			phrase[i] += "ay"
			continue
		} else {

		}

	}
	return strings.Join(phrase, "")
}
