package main

import "fmt"

// start by: go run hello.go

func add(x int, y int) int {
    return x + y
}

func main() {
	fmt.Println(add(2,3))
    fmt.Println("Hello, World")
}