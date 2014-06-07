package main

import "fmt"

func sender(c chan int) {
    c <- 2
    c <- 3
    c <- 4
    close(c)
}

func sender2(c chan int) {
    c <- 5
    c <- 6
    c <- 9
    c <- 10
    c <- 15
}

func printer(c chan int) {
    for i := range c {
        fmt.Println("message received:", i)
    }
}

func main() {
    c := make(chan int)

    go sender2(c)
    go sender(c)
    printer(c)
}