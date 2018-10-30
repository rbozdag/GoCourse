package main

import (
	"fmt"
)

func main() {
	a := 10
	b := &a

	var c *int
	c = &a

	var d *int
	d = c

	fmt.Printf("a: %d, type: %T\n", a, a)
	fmt.Printf("b: %d, type: %T\n", *b, b)
	fmt.Printf("c: %d, type: %T\n", *c, *c)
	fmt.Printf("d: %d, type: %T\n", *d, d)

	*c = 11
	fmt.Printf("after changed) a: %d, type: %T\n", a, a)
	fmt.Printf("after changed) b: %d, type: %T\n", *b, b)
	fmt.Printf("after changed) c: %d, type: %T\n", *c, *c)
	fmt.Printf("after changed) d: %d, type: %T\n", *d, d)

}
