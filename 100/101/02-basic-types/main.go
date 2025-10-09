package main

import "fmt"

func main() {
	a := 1
	b := 2
	c := 3
	d := 4
	e := 1
	satisfyTheCompiler(a, b, c, d, e)
}

func satisfyTheCompiler(a int, b int64, c float64, d string, e bool) {
	// %d prints integers, you can also use %x for hex, %o for octal, etc.
	fmt.Printf("a = %d\n", a)
	fmt.Printf("b = %d\n", b)
	// %f prints floating point numbers, you can specify precision like %.2f for 2 decimal places.
	fmt.Printf("c = %f\n", c)
	// %s prints the string, there's also %q which adds quotes around the string.
	fmt.Printf("d = %s\n", d)
	// %v prints the value in a default format.
	fmt.Printf("e = %v\n", e)
}
