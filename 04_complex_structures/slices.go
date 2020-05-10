package main

import "fmt"

func main() {
	var mySlice []int = make([]int, 5, 10)
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))

	fruitArray := [5]string{"watermelon", "grapes", "strawberry", "cookies", "cake"}
	fmt.Println(fruitArray)

	var splicedFruit []string = fruitArray[1:3]
	fmt.Println("This is splicedFruit: ", splicedFruit)

	// Append
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice2)

	originalSlice := []int{1, 2, 3}
	destination := make([]int, len(originalSlice))
	fmt.Println("Before copy:", originalSlice, destination)

	mysteryValue := copy(destination, originalSlice)
	fmt.Println("The mystery value: ", mysteryValue) //3
	fmt.Println("After copy:", originalSlice, destination)

}
