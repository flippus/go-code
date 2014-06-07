package main

import "fmt"

func sender(c chan int) {
    c <- 2
    c <- 3
    close(c)
}

func printer(c chan int) {
    for i := range c {
        fmt.Println("message received ", i)
    }
}

func main() {
    c := make(chan int)

    go sender(c)
    printer(c)
}