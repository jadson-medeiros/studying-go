package main

import "fmt"

type CheckingAccount struct {
	owner         string
	agencyNumber  int
	accountNumber int
	balance       float64
}

func (c *CheckingAccount) Withdraw(withdrawalAmount float64) string {
	canDraw := withdrawalAmount > 0 && withdrawalAmount <= c.balance
	if canDraw {
		c.balance -= withdrawalAmount
		return "Successful withdrawal"
	} else {
		return "Insufficient funds"
	}
}

func main() {
	jasonAccount := CheckingAccount{}
	jasonAccount.owner = "Jason"
	jasonAccount.balance = 376.32

	fmt.Println(jasonAccount.balance)

	fmt.Println(jasonAccount.Withdraw(100))
	fmt.Println(jasonAccount.balance)
}
