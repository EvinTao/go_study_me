package main

import (
	"fmt"
	"log"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	//Counter
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
			log.Println("----")
		}
		close(naturals)
	}()

	//squarer
	go func() {
		//for {
		//	x, ok := <-naturals
		//	if !ok {
		//		break // channel was closed and drained
		//	}
		//	squares <- x * x
		//}
		//close(squares)
		for x := range naturals {
			squares <- x * x
			log.Println("+++++++")
		}
		close(squares)
	}()

	//Pointer (in main goroutine)
	//for {
	//	fmt.Println(<-squares)
	//}
	for x := range squares {
		fmt.Println(x)
	}

}
