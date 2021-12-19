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

func forExampleLikeC(numbers ...int) {
	for i := 0; i < len(numbers); i++ {
		fmt.Println(i, numbers[i])
	}
}

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	forExampleLikeC(1, 2, 3, 4, 5, 6)
	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
}
