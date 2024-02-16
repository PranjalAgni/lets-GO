package main

import "fmt"

func makeAge(c chan int, age int) {
	c <- age * 2
}

func main() {
	resultChan := make(chan int)
	go makeAge(resultChan, 10)
	go makeAge(resultChan, 12)

	modifiedAge1 := <-resultChan
	modifiedAge2 := <-resultChan

	fmt.Println(modifiedAge1, modifiedAge2)
	// modifiedAge3, err := <-resultChan
	// fmt.Println(modifiedAge3, err)

}
