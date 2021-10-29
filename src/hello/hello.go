package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Hello world with Go")

	showIntroduction()

	for {
		showMenu()

		command := getCommand()

		switch command {
		case 1:
			startMonitor()
		case 2:
			fmt.Println("Showing Logs...")
		case 0:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("I don't know this command.")
			os.Exit(-1)
		}
	}
}

func showIntroduction() {
	name := "GitHub"
	version := 1.0

	fmt.Println("Hello, sr.", name)
	fmt.Println("This program is available to version", version)
}

func showMenu() {
	fmt.Println("1- Start Monitoring")
	fmt.Println("2- Show Logs")
	fmt.Println("0- Exit")
}

func getCommand() int {

	var command int
	fmt.Scan(&command)
	fmt.Println("The chosen command was", command)

	return command
}

func startMonitor() {
	fmt.Println("Monitoring...")

	site := "https://www.github.com"

	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "was loaded with success!")
	} else {
		fmt.Println("Site:", site, "has issues. Status Code:", resp.StatusCode)
	}
}
