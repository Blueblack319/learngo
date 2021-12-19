package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func multiArgs(strings ...string) []string {
	return strings
}

func main() {
	names := multiArgs("James", "Arts", "huddy")
	fmt.Println(names)
}
