package main

import "fmt"

func main() {
	var s1 string
	var s2 string

	numOfArgs, err := fmt.Scan(&s1, &s2)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Number of args:", numOfArgs, "String 1:", s1, "String 2:", s2)
}
