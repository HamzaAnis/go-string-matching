package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	code, _ := ioutil.ReadFile("IccUtils.java")
	codeStr := string(code)
	// rdr := bufio.NewReader(os.Stdin)
	// fmt.Println("Please enter a string for matching")
	// p, _ := rdr.ReadString('\n')

	findPackage(codeStr)
}
func findPackage(text string) {

	pckg := "package"
	// pckgName := make([]string, 0)

	pckgName := ""
	fmt.Println(len(text))
	for i := 0; i < len(text); i++ {
		j := 0
		for j = 0; j < len(pckg); j++ {
			if string(text[i+j]) != string(pckg[j]) {
				break
			}
		}
		if j == len(pckg) {
			fmt.Println("It is found at", i, string(text[i:i+len(pckg)]))
			for text[i] != '\n' {
				pckgName += string(text[i])
				i++
			}
		}

	}
	fmt.Println(pckgName)
}
