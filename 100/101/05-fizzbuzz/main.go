package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/a-h/go-workshop/100/101/05-fizzbuzz/fizzbuzz"
)

func main() {
	cmd := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	n := cmd.Int("n", 0, "Number to check whether it's Fizz or Buzz")
	cmd.Parse(os.Args[1:])

	fb := fizzbuzz.Check(*n)
	fmt.Printf("n=%d, fizzbuzz=%s\n", *n, fb)
}
