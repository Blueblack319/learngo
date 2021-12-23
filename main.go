package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	people := [2]string{"hyang", "gwang"}
	for _, person := range people {
		go isSexy(person, c)
	}
	resultOne := <-c
	resultTwo := <-c
	fmt.Println(resultOne)
	fmt.Println(resultTwo)
}

func isSexy(name string, c chan string) {
	time.Sleep(time.Second * 2)
	c <- name + " is sexy"
}
