package main

import (
	"strconv"
	"strings"
)

func sortSentence(s string) string {
	words := strings.Split(s, " ")
	wordsArray := make([]string, len(words))
	for _, word := range words {
		wordString := word[len(word)-1:]
		index, _ := strconv.Atoi(wordString)
		index -= 1
		wordsArray[index] = word[:len(word)-1]
	}

	return strings.Join(wordsArray, " ")

}
