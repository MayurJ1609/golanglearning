package main

import (
	"fmt"
	"time"
)

func main() {
	go greeter("Hello")
	greeter("World")
}

func greeter(s string) {

	for i := 0; i < 6; i++ {
		time.Sleep(5 * time.Second)
		fmt.Println(s)
	}
}
