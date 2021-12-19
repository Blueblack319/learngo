package main

import (
	"fmt"
	"log"

	"github.com/crazybirdz/learngo/myDict"
)

func main() {
	dictionary := myDict.Dictionary{"first": "First word"}
	baseWord := "hello"
	err := dictionary.Add(baseWord, "Greeting")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(dictionary)
	_, err2 := dictionary.Search(baseWord)
	if err2 != nil {
		log.Fatalln(err)
	}
	err3 := dictionary.Update(baseWord, "Hi")
	if err3 != nil {
		log.Fatalln(err3)
	}
	fmt.Println(dictionary)
	// err4 := dictionary.Update("jiojvcxz", "jbviofaj")
	// if err4 != nil {
	// 	log.Fatalln(err4)
	// }
	dictionary.Delete(baseWord)
	fmt.Println(dictionary)
}
