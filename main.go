package main

import (
	"fmt"
	"log"

	"github.com/crazybirdz/learngo/myDict"
)

func main() {
	dictionary := myDict.Dictionary{"first": "First word"}
	err := dictionary.Add("hello", "Greeting")
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println(dictionary)
	}
	err2 := dictionary.Add("hello", "Hi")
	if err2 != nil {
		log.Fatalln(err2)
	} else {
		fmt.Println(dictionary)
	}
}
