package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	code, _ := ioutil.ReadFile("IccUtils.java")
	fmt.Println(len(string(code)))
	fmt.Println(len(code))
	fmt.Println("Hello World")
}
