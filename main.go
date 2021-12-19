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

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 20 {
		return false
	}
	return true
}

func canIDrinkSwitch(age int) bool {
	switch koreanAge := age + 2; koreanAge {
	case 18:
		return false
	case 19:
		return false
	}
	return true
}

func main() {
	fmt.Println(canIDrinkSwitch(18))
}
