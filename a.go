package main

import "fmt"

func main() {
	fmt.Println("Welcome to the  world!")
	var x string
	fmt.Printf("Enter your name:  ")
	fmt.Scan(&x)
	fmt.Println("Hi", x, "! How are you?\n Welcome.....")
	var age int
	fmt.Printf("Enter your age: ")
	fmt.Scan(&age)
	if age < 10 {
		fmt.Print("You arent eligible\n")
	} else {
		fmt.Print("Yay! You can play\n")
		return
	}

}
