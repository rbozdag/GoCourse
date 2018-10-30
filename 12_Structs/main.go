package main

import (
	"fmt"
)

type User struct {
	firstName string
	lastName  string
	city      string
	age       int
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
	userPointer.age *= 2
	fmt.Println(user, " --> original variable")
	fmt.Println(*userPointer, " -> pointer")

}
