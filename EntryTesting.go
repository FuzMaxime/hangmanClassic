package Hangman

import "fmt"

func EntryTest(word string, lWordFull []string, listEmpty []string, letter string, listLetterEntered []string) int {
	test := 0
	count := 0
	/*------- test if letter is a letter ------*/
	if letter < "a" || letter > "z" {
		fmt.Printf("\nYou didn't type a lowercase letter !\n")
		test = 1
		return test
	}
	/*------- test if letter is already typed ------*/
	for i := 0; i < len(listLetterEntered); i++ {
		if letter == listLetterEntered[i] {
			fmt.Printf("\nYou have already typed the letter !\n")
			test = 1
			return test
		}
	}
	/*--------check if the letter is in the word ---------*/
	for position := 0; position < len(lWordFull); position++ {
		if string(letter) == string(lWordFull[position]) {
			listEmpty[position] = string(letter)
			count += 1
		}
	}
	/*------ display if the letter is not in the word ----*/
	if count == 0 {
		fmt.Println("The letter :", string(letter), " is not in this word !")
		test = 2
	}
	return test
}
