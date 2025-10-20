package main

import (
	"fmt"

	"github.com/ppahari/go-project/pointer-struct/task"
)

func main() {

	email := "partha.pahari@gmail.com"

	user := task.User{
		Name:  "Partha",
		Email: &email,
	}
	fmt.Println(user.Name)
	fmt.Println("Email:", *user.Email)

	emp := task.Employee{Name: "Partha", Level: 1}
	task.Promote(&emp)

	fmt.Println(emp)

	//Struct contaning Pointer to another struct
	addr := &task.Address{City: "Bangalore", Country: "India"}

	person := task.Person{
		Name:    "Partha",
		Address: addr,
	}

	fmt.Println(person.Name, "lives in", person.Address.City)

	//update the nested strcut through pointer
	person.Address.City = "Kolkata"

	fmt.Println("Moved to:", person.Address.City)

}
