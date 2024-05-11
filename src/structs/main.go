package main

import "fmt"

type Adreess struct {
	Street string
	Number int
	City string
	State string
}

type Client struct {
	Name string
	Age int
	Status bool
	Adreess
}

func (c* Client) Desactivate() {
	c.Status = false
	fmt.Printf("The client %s is now inactive", c.Name)
}	

func main(){
	jadson := Client{
		Name: "Jadson",
		Age: 27,
		Status: true,
	}

	jadson.Desactivate()
	fmt.Println(jadson)
}