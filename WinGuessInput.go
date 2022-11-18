package Hangman

import "fmt"

func Input_letterOrWord(listLetterEntered []string) string {
	fmt.Print("Enter a letter or the word : ")
	var letter string
	fmt.Scanln(&letter)
	return letter
}

func List_word(lWordFull []string, letters []string) []string {
	lWordEmpty := make([]string, len(lWordFull))
	for iEmpty := 0; iEmpty < len(lWordEmpty); iEmpty++ {
		lWordEmpty[iEmpty] = "_"
	}
	for letter := 0; letter < len(letters); letter++ {
		for index := 0; index < len(lWordFull); index++ {
			if letters[letter] == lWordFull[index] {
				lWordEmpty[index] = letters[letter]
			}
		}
	}
	return lWordEmpty
}

func ListWordFull(word string) []string {
	listTransition := []string{}
	for _, letterWord := range word {
		listTransition = append(listTransition, string(letterWord))
	}
	lWordFull := []string{}
	wordFull := ""
	for i := 0; i < len(listTransition)-1; i++ {
		wordFull += string(listTransition[i])
	}
	/* wordFull = word */
	for _, j := range wordFull {
		lWordFull = append(lWordFull, string(j))
	}
	/* lWordFull = [w o r d] */
	return lWordFull
}

func Win(listTest []string) bool {
	reply := true
	for v := 0; v < len(listTest); v++ {
		if listTest[v] < "a" || listTest[v] > "z" {
			reply = false
		}
	}
	return reply
}

func Guess(lWordFull []string, word string) bool {
	reply := true
	lLetter := []string{}
	for _, v := range word {
		lLetter = append(lLetter, string(v))
	}
	for position := 0; position < len(lLetter); position++ {
		if lLetter[position] != lWordFull[position] {
			return false
		}
	}
	return reply
}
