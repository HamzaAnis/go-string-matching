package main

import "io/ioutil"

func main() {
	fileName := []string{"IccUtils.java"}
	for _, val := range fileName {
		code, _ := ioutil.ReadFile(val)
		codeStr := string(code)
		// word, _ := findMatch(codeStr, "package", '\n')
	}
}
func findMatch(text string, toFind string, terminator int) (string, int) {
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
		// if j == len(toFind) {
		// 	fmt.Println("Found")
		// 	j = lps[j-1]
		// }
	}

	i = 0
	for i < len(text) {
		if string(toFind[i]) == string(toFind[l]) {
			j++
			i++
		} else {
			if l != 0 {
				l = lps[l-1]
			} else {
				lps[i] = l
				i++
			}
		}
		// if j == len(toFind) {
		// 	fmt.Println("Found")
		// 	j = lps[j-1]
		// }
	}

}
