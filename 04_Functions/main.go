package main

import "fmt"

func main() {
	fmt.Println(greeting("Rahmi"))
}

//private access
func greeting(name string) string {
	return "Hello " + name
}
