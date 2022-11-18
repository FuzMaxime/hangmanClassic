package Hangman

import (
	"io/ioutil"
	"strings"
)

func Index(list []string) []int {
	/* Convert list string to rune */
	str := ""
	for _, strList := range list {
		str += strList
	}
	index := []int{}
	listRune := []rune(str)
	/* Convert list rune to index of the letter in standard.txt */
	for _, indexList := range listRune {
		if indexList == '_' {
			index = append(index, int(1))
		} else {
			index = append(index, int(rune(indexList)-97+2))
		}
	}
	return index
}

func WordAscii(list []int) string {
	file, _ := ioutil.ReadFile("../standard.txt")
	fileAscii := string(file)
	asciiWord := ""
	lignes := []string{}
	listLignes := make([]string, 8)
	letters := strings.Split(string(fileAscii), ",,")
	for iLigne := 0; iLigne < 8; iLigne++ {
		for index := 0; index < len(list); index++ {
			lignes = strings.Split(letters[list[index]], "\n")
			for _, iCaracter := range lignes[iLigne] {
				if iCaracter != rune(13) && iCaracter != rune(10) {
					listLignes[iLigne] += string(iCaracter)
				}
			}
		}
		listLignes[iLigne] += string("\n")
	}
	for iListLigne := 0; iListLigne < len(listLignes); iListLigne++ {
		asciiWord += listLignes[iListLigne]
	}
	return asciiWord
}
