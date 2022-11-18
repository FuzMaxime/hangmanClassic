package main

import (
	"Hangman"
	"io/ioutil"
)

/* 	test début
forme code*/

func main() {
	file, _ := ioutil.ReadFile("../words.txt") //read the txt
	word := Hangman.RandomWord(file)           //import the function of randomise
	letters := Hangman.RandomLetters(word)
	lWordFull := Hangman.ListWordFull(word)
	Hangman.GameDisplay(word, lWordFull, letters)
}
