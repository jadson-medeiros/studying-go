package main

import "fmt"

func main() {
	fmt.Println("Hello world with Go")

	name := "GitHub"
	version := 1.0

	fmt.Println("Hello, sr.", name)
	fmt.Println("This program is available to version", version)

	fmt.Println("1- Start Monitoring")
	fmt.Println("2- Show Logs")
	fmt.Println("0- Exit")

	var command int
	fmt.Scan(&command)
	fmt.Println("The chosen command was", command)
}
