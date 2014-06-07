package main

import "fmt"

func sender(r, s chan string) {
    r <- "Hello"

    for i := range s {
        fmt.Println("message sent back:", i)
        close(r)
    }
}

func printer(s, r chan string) {
    for i := range r {
        fmt.Println("message received:", i)
        s <- i
        close(s)
    }
}

func main() {
    senderChannel := make(chan string)
    receiverChannel := make(chan string)

    go sender(receiverChannel, senderChannel)
    printer(senderChannel, receiverChannel)
}