package main

import (
	"fmt"
)

func main() {
	cs := make(chan int)//remove the receive only channel

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}
