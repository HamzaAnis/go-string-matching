package main

import (
	"io/ioutil"
)

func main() {
	fileName := []string{"IccUtils.java"}
	for _, val := range fileName {
		code, _ := ioutil.ReadFile(val)
		codeStr := string(code)
		findMatch(codeStr, "package", '\n')
	}
}
func findMatch(text string, toFind string, terminator int) (string, int) {
	word := ""
	count := 0
	lps := make([]int, len(toFind))
	j := 0

	l := 0
	i := 1
	lps[0] = 0

	for i < len(toFind) {
		if string(toFind[i]) == string(toFind[l]) {
			j++
			i++
			lps[i] = l
		} else {
			if l != 0 {
				l = lps[l-1]
			} else {
				lps[i] = l
				i++
			}
		}

	}

	i = 0
	j = 0
	for i < len(text) {
		if string(toFind[j]) == string(text[i]) {
			j++
			i++
		}
		if j == len(toFind) {
			for text[i] != byte(terminator) {
				word += string(text[i])
				i++
				count++
			}
			j = lps[j-1]
		} else if i < len(text) && string(toFind[j]) != string(text[i]) {
			if j != 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}
	return word, count
}
