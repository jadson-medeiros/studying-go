package main

import (
	"encoding/json"
	"os"
)

type Account struct {
	Number int `json:"number"`
	Money int `json:"money"`
}

func main() {
	account := Account{Number:1, Money: 100}
	res, err := json.Marshal(account)

	if err != nil {
		println(err)
	}

	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(account)
	if err != nil {
		println(err)
	}

	jsonPure := []byte(`{"number":1, "money":2000}`)
	var account2 Account
	err = json.Unmarshal(jsonPure, &account2)
	if err != nil {
		println(err)
	}
	println(account2.Money)
}