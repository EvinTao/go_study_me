package main

import (
	"fmt"
	"time"
)

func counter(out chan<- int) {
	for x := 1; x < 10; x++ {
		out <- x
		time.Sleep(1 * time.Second)
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	naturals := make(chan int)
	squarers := make(chan int)

	go counter(naturals)
	go squarer(squarers, naturals)
	printer(squarers)
}
