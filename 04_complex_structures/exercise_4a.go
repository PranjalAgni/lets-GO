package main

import "fmt"

func average(num1, num2, num3 float64) float64 {
	return (num1 + num2 + num3) / 3
}

func main() {
	result := average(5, 5, 5)
	fmt.Println("Result: ", result)
}
