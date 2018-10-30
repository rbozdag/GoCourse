package main

import (
	"fmt"
)

func main() {
	const name string = "Aylin"
	var age = 2
	birthDate := "05/10/2016"

	fmt.Println("She is ", name, " and ", age, " years old")
	fmt.Printf("type of variable name %T \n", name)
	fmt.Println("birthDate", birthDate)
}

//Main Types
//  string
//  bool
//  int
//  uint
//  byte -alias for uint8
//  rune
//  complex64, complex128
