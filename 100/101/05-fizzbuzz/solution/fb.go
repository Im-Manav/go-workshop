package fizzbuzz

import "fmt"

func Check(n int) (result string) {
	if n%3 == 0 && n%5 == 0 {
		return "FizzBuzz"
	}
	if n%3 == 0 {
		return "Fizz"
	}
	if n%5 == 0 {
		return "Buzz"
	}
	return fmt.Sprintf("%d", n)
}
