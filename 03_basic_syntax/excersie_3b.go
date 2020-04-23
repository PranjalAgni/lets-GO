package main

import (
	"fmt"
	"strings"
)

func main() {
	var name, city, weather string
	var years int
	fmt.Println("Enter your name")
	fmt.Scan(&name, "%s")
	fmt.Println("What city they live in")
	fmt.Scan(&city, "%s")
	fmt.Println("How long they have lived there")
	fmt.Scan(&years, "%d")
	fmt.Println("If the weather is nice")
	fmt.Scan(&weather, "%s")

	weather = strings.ToLower(weather)
	weatherStatus := false

	if weather == "yes" {
		weatherStatus = true
	}

	fmt.Printf("Hi! My name is %s. I have lived in %s for %d years. They say the weather is amazing, which is %t", name, city, years, weatherStatus)
}
