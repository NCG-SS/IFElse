package main

import (
	"fmt"
)

func main() {
	var username string
	fmt.Println("Enter your username: ")
	fmt.Scanln(&username)

	if username == "Robin" || username == "John" {
		fmt.Println("Welcome!")
	} else if username == "Admin" {
		fmt.Println("Welcome Admin!")
	} else {
		fmt.Println("Merry men")
	}
}
