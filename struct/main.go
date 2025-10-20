package main

import "fmt"

type person struct {
	name string
	age  int
	pet  string
}

func main() {
	var fred person
	bob := person{}
	julia := person{
		name: "Julia",
		age:  25,
		pet:  "cat",
	}
	fmt.Println(fred)
	fmt.Println(bob)

	fmt.Println(julia)
}
