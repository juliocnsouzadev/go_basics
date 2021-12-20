package main

import (
	"fmt"

	"github.com/juliocnsouzadev/go_basics/bar"
	"github.com/juliocnsouzadev/go_basics/foo"
	"github.com/juliocnsouzadev/go_basics/math"
)

func main() {
	fooBar()
	calcuations()
}

func calcuations() {
	add := math.Add(1, 2)
	sub := math.Sub(1, 2)
	mult := math.Multiply(1, 2)
	div := math.Divide(1, 0)
	fmt.Printf("Add: %d, Sub: %d, Mult: %d, Div: %d", add, sub, mult, div)
}

func fooBar() {
	foo.Foo()
	bar.Bar()
}
