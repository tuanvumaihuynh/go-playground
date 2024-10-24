# Buffered Channel vs Unbuffered Channel

## Buffered Channel

- A buffered channel is a channel that has a buffer.
- A buffered channel can store a limited number of values without a corresponding receiver for those values.
- Communication over a buffered channel does not block if the buffer is not full (Asynchronous Communication).

```go
package main

import "fmt"

func main() {
    // Create a buffered channel with a capacity of 2
    ch := make(chan int, 2)

    // Send two values to the channel
    ch <- 1
    ch <- 2

    // Receive the values from the channel
    fmt.Println(<-ch)
    fmt.Println(<-ch)
}
```

## Unbuffered Channel

- An unbuffered channel is a channel that has no buffer.
- An unbuffered channel can store only one value at a time.
- Communication over an unbuffered channel blocks until the sender and receiver are ready (Synchronous Communication).

```go
package main

import "fmt"

func main() {
    // Create an unbuffered channel
    ch := make(chan int)

    // Send a value to the channel
    go func() {
        ch <- 1
    }()

    // Receive the value from the channel
    fmt.Println(<-ch)
}
```