package main

import "fmt"

func main() {
	fmt.Println(greeting("Rahmi"))
}

func greeting(name string) string {
	return "Hello " + name
}
