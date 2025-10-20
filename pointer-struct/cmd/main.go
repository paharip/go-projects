package main

import (
	"fmt"
	"github.com/ppahari/go-project/pointer-struct/task"
)

func main() {

	email := "partha.pahari@gmail.com"

	user := task.User{
		Name: "Partha",
		Email: &email,
	}
	fmt.Println(user.Name)
	fmt.Println("Email:", *user.Email)
}
