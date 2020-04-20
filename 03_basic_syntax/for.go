package main

import "fmt"

func main() {

	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}

	fmt.Println("Sum is: ", sum)

	var sentence string = "Hello world this"

	for idx, letter := range sentence {
		if idx%2 == 0 {
			fmt.Print(string(letter) + " ")
		}
	}
}
