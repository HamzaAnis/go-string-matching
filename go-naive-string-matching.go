package main

import (
	"io/ioutil"

	"github.com/fatih/color"
)

func main() {
	fileName := []string{"IccUtils.java", "EncodeException.java", "GsmAlphabet.java"}
	for _, val := range fileName {
		code, _ := ioutil.ReadFile(val)
		codeStr := string(code)
		word, _ := findMatch(codeStr, "package", '\n')
		color.Blue("Package name" + word)
		word, _ = findMatch(codeStr, "class", '{')
		color.Green("\t\t\tClass name\n" + word[:len(word)-2] + "\n")
		word, _ = findMatch(codeStr, "public", '{')
		color.Yellow("\t\t\tMethods\n", word)
		word, _ = findMatch(codeStr, "private", '{')
		color.Yellow(word)

		_, count := findMatch(codeStr, "for", '\n')
		color.Blue(" Loops are %v ", count)
	}
}
func findMatch(text string, toFind string, terminator int) (string, int) {
	pckgName := ""
	count := 0

	for i := 0; i < len(text); i++ {
		j := 0
		for j = 0; j < len(toFind); j++ {
			if string(text[i+j]) != string(toFind[j]) {
				break
			}
		}
		if j == len(toFind) {
			for text[i] != byte(terminator) {
				pckgName += string(text[i])
				i++
			}
			pckgName += "\n"
			count++
		}
	}
	return pckgName, count
}
