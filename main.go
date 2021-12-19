package main

import (
	"fmt"
	"time"
)

func main() {
	go whoIsSexy("Hoon")
	whoIsSexy("Aiden")
}

func whoIsSexy(name string) {
	for i := 0; i < 10; i++ {
		fmt.Println(name, "is sexy.", i)
		time.Sleep(time.Second)
	}
}
