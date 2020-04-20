// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func someFunction() error {
// 	return errors.New("Some error")
// }

// func main() {
// 	var someVar int = 33
// 	if someVar > 30 {
// 		fmt.Println("This is ", someVar)
// 	}

// 	someVar = 100

// 	if someVar > 100 {
// 		fmt.Println("Greater than 100")
// 	} else if someVar == 100 {
// 		fmt.Println("Equals 100")
// 	} else {
// 		fmt.Println("Less then 100")
// 	}

// 	err := someFunction()

// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// }
