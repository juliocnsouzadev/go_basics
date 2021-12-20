package math

import (
	"fmt"
)

func Add(x, y int) int {
	return x + y
}

func Sub(x, y int) int {
	return x - y
}

func Multiply(x, y int) int {
	return x * y
}

func Divide(x, y int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic occurred:", err)
		}
	}()
	return x / y
}

func Mod(x, y int) int {
	return x % y
}
