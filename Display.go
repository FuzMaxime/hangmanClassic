package Hangman

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func GameDisplay(word string, lWordFull []string, letters []string) {
	listEmpty := List_word(lWordFull, letters)
	fmt.Printf("----------------------------\nThe word to find is : ---------------------------- %v\n", WordAscii(Index(listEmpty)))
	Attempt := 10
	nbrOfAttempt := 10
	listLetterEntered := []string{}
	for Attempt > 0 {
		check := true
		fmt.Printf("You entered his letters %v\n", listLetterEntered)
		letter := Input_letterOrWord(listLetterEntered)
		if len(letter) > 1 {
			/*------------------- display if GUESS ----------------*/
			if Guess(lWordFull, letter) {
				fmt.Printf(" ----------------------------The word is : ---------------------------- %v\n", WordAscii(Index(lWordFull)))
				reply := ""
				data, err := ioutil.ReadFile("../win.txt")
				if err != nil {
					fmt.Println("Error of reading file", err)
					return
				}
				reply += string(data)
				fmt.Println(reply)
				return
			} else {
				if nbrOfAttempt == 1 {
					nbrOfAttempt -= 1
				} else {
					nbrOfAttempt -= 2
				}
			}
		} else {
			/*------ if the letter is not in the word ----*/
			if EntryTest(word, lWordFull, listEmpty, letter, listLetterEntered) == 2 {
				listLetterEntered = append(listLetterEntered, string(letter))
				nbrOfAttempt -= 1
				/*------ if the letter is in the word ----*/
			} else if EntryTest(word, lWordFull, listEmpty, letter, listLetterEntered) == 0 {
				listLetterEntered = append(listLetterEntered, string(letter))
			} else {
				check = false
			}
		}
		if check {
			/*------------------- display if WIN ----------------*/
			if Win(listEmpty) {
				fmt.Printf(" ----------------------------The word is : ---------------------------- %v\n", WordAscii(Index(listEmpty)))
				reply := ""
				data, err := ioutil.ReadFile("../win.txt")
				if err != nil {
					fmt.Println("Error of reading file", err)
					return
				}
				reply += string(data)
				fmt.Println(reply)
				return
			}
			/*------------------- display jose ----------------*/
			if nbrOfAttempt < 10 {
				reply := ""
				nbrTxt := int((nbrOfAttempt - 10) * -1)
				txt := "../"
				txt += strconv.Itoa(nbrTxt)
				txt += ".txt"
				data, err := ioutil.ReadFile(txt)
				if err != nil {
					fmt.Println("Error of reading file", err)
					return
				}
				reply += string(data)
				fmt.Println(reply)
			}
			/*----------------- display word --------------*/
			if nbrOfAttempt == 0 {
				fmt.Printf("The word was : %v\n", word)
			}
			if nbrOfAttempt > 0 {
				fmt.Printf("You still have %v attempts !\n", nbrOfAttempt)
				fmt.Printf("	%v\n----------------------------\n", WordAscii(Index(listEmpty)))
			} else {
				return
			}
		}
	}
}
