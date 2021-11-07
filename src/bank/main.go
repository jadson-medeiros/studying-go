package main

import "fmt"

type CheckingAccount struct {
	owner         string
	agencyNumber  int
	accountNumber int
	balance       float64
}

func main() {
	rhysAccount := CheckingAccount{"Rhys",
		672,
		456321,
		145.98,
	}

	fmt.Println(rhysAccount)
}
