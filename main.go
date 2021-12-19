package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done.")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func multiArgs(strings ...string) []string {
	return strings
}

func main() {
	leng, _ := lenAndUpper("James")
	fmt.Println(leng)
}
