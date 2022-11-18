package Hangman

import (
	"math/rand"
	"strings"
	"time"
)

func RandomWord(file []byte) string {
	words := strings.Split(string(file), "\n")
	rand.Seed(time.Now().UnixNano())
	word := rand.Intn(len(words))
	return words[word]
}

func RandomLetters(word string) []string {
	listWord := []string{}
	listReply := make([]string, len(word)/2-1)
	for _, strWord := range word {
		listWord = append(listWord, string(strWord))
	}
	if len(listWord) < 4 {
		listReply = append(listReply, "") // No letters because the word lenth is < 4
		return listReply
	} else {
		for i := 0; i < len(listWord)/2-1; i++ {
			rand.Seed(time.Now().UnixNano())
			letter := rand.Intn(len(listWord))
			letterAppend := listWord[letter]
			test := 1 // false
			for j := 0; j < len(listReply); j++ {
				if letterAppend == listReply[j] {
					test = 0 // true
				}
			}
			if test == 1 {
				listReply = append(listReply, letterAppend)
			} else {
				i -= 1
			}
		}
	}
	return listReply
}
