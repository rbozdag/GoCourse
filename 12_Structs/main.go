package main

import (
	"fmt"
	"strconv"
)

type User struct {
	firstName string
	lastName  string
	city      string
	age       int
}

func (u User) greet() string {
	return "Hello I am " + u.firstName + " " + u.lastName + " and I am " + strconv.Itoa(u.age) + " years old."
}

func (u *User) updateAge(age int) {
	u.age = age
}

func main() {
	//descriptive way
	user := User{firstName: "Rahmi", lastName: "Bozdağ", city: "İstanbul", age: 30}
	fmt.Println(user, " --> descriptive way")

	fmt.Println("----------------------------------------")

	//short way
	userShort := User{"Rahmi", "Bozdağ", "İstanbul", 30}
	fmt.Println(userShort, " --> short way")

	fmt.Println("----------------------------------------")

	//assign with copy
	userCopy := user
	userCopy.age++
	fmt.Println(user, " --> original variable")
	fmt.Println(userCopy, " -> copy of variable")

	fmt.Println("----------------------------------------")

	//pointer
	userPointer := &user
	userPointer.age++
	fmt.Println(user, " --> original variable")
	fmt.Println(*userPointer, " -> pointer")

	fmt.Println("----------------------------------------")

	// extensions
	fmt.Println(user.greet(), " --> greet")

	user.updateAge(35)
	fmt.Println(user.greet(), " --> updateAge after greet")
}
