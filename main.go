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
	/**
	 * There are four instances for pig latin:
	 * Word starts with a vowel, just add "way", artist -> artistway
	 * Word starts with a consonant, move letter to end and add "ay", food -> oodfay
	 * Word starts with consonant cluster, move the cluster to the  end and add "ay", swimming -> immingsway
	 * Word has no vowel, impossible
	 */
	for i := range phrase {
		// If it starts with a vowel, just add "way"
		if isVowel(string(phrase[i][0])) {
			phrase[i] += "way"
			continue
		}
		// If its a small word and doesnt start with a vowel, just add "ay"
		if len(phrase[i]) < 2 {
			phrase[i] += "ay"
			continue
		}
		// Ok so what this can do is loop until the first letter of the word is not a vowel
		for {
			if !isVowel(string(phrase[i][0])) {
				var tmp string = string(phrase[i][0])
				phrase[i] = phrase[i][1:]
				phrase[i] += tmp

				fmt.Println(phrase[i])

			} else {
				phrase[i] += "ay"
				break
			}
		}
		fmt.Println(phrase[i])
	}
	return strings.ToLower(strings.Join(phrase, " "))
}
