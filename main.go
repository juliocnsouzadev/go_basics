package main

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/juliocnsouzadev/go_basics/bar"
	"github.com/juliocnsouzadev/go_basics/foo"
	"github.com/juliocnsouzadev/go_basics/math"
)

func main() {
	fooBar()
	breakLine()
	calcuations()
	breakLine()
	{
		checkType(1)
		checkType(1.0)
		checkType(false)
		checkType("cool")
		breakLine()
	}
	casting()
	breakLine()
}

func breakLine() {
	fmt.Println("")
}

func calcuations() {
	fmt.Println("Enter 2 integer numbers to perform calcuations: ")

	var a, b int
	_, err := fmt.Scanf("%d %d", &a, &b)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	add := math.Add(a, b)
	sub := math.Sub(a, b)
	mult := math.Multiply(a, b)
	div := math.Divide(a, b)
	fmt.Printf("Add: %d, Sub: %d, Mult: %d, Div: %d\n", add, sub, mult, div)
}

func fooBar() {
	foo.Foo()
	bar.Bar()
}

func checkType(i interface{}) {
	fmt.Printf("value: %v, type %T\n", i, i)
	fmt.Println("Type: ", reflect.TypeOf(i))
}

func casting() {
	var f float64 = 3.4
	var i int = int(f)
	fmt.Printf("float %f to int %d\n", f, i)

	var s string = strconv.Itoa(i)
	fmt.Printf("int %d to string %s\n", i, s)
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Printf("string %s to int %d\n", s, i)
}
