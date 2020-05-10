package main

import "fmt"

func main() {
	// var userEmails map[int]string
	userEmails := make(map[int]string)
	userEmails[1] = "pranjal@io.com"
	userEmails[2] = "anay@io.com"
	fmt.Println(userEmails)

	firstName, ok := userEmails[1]
	fmt.Println(firstName, ok)

	if _, ok := userEmails[1]; ok {
		fmt.Println("Yes it exists")
	} else {
		fmt.Println("I don't know what you want from me")
	}

	for k, v := range userEmails {
		fmt.Printf("%d has email %s \n", k, v)
	}
}
