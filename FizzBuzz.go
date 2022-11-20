package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	go FizzBuzz()

	fmt.Println("Press enter to exit")
	fmt.Scanln()

}

func FizzBuzz() {

	var n int
	for {
		n = rand.Intn(100) + 1

		switch {

		case n%15 == 0:
			fmt.Println("FizzBuzz")

		case n%3 == 0:
			fmt.Println("Fizz")

		case n%5 == 0:
			fmt.Println("Buzz")

		default:
			fmt.Println(n)
		}

		time.Sleep(1000 * time.Millisecond)
	}
}
