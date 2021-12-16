package main

import (
	"fmt"

	"Github/julio-repos/go_basics/app/network"
)

func main() {
	network.DoGet("'https://jsonplaceholder.typicode.com/posts/1")
	fmt.Println("Kickstart Go")
}
