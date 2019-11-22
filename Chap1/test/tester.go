package test

import (
	"fmt"
)

// Adder  is an adder function with two parameters and one return value
func Adder(one int, two int) int {
	fmt.Println(one + two)
	return (one + two)
}
