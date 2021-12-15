package main

import "fmt"

type User struct {
	name  string
	email string
	age   int
}

func (user *User) changeName(name string) {
	user.name = name
}

func main() {
	user := &User{
		name:  "taro",
		email: "taro@example.com",
		age:   33,
	}

	fmt.Println(user.name)

	var changedName string
	fmt.Print("name? > ")
	fmt.Scan(&changedName)
	user.changeName(changedName)
	fmt.Println(user.name)
}
