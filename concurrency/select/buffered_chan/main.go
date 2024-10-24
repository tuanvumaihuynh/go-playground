package main

import "fmt"

func main() {
	ch := make(chan string, 3)
	chars := []string{"a", "b", "c"}

	for _, s := range chars {
		// select {
		// case ch <- s:
		// }
		ch <- s
	}

	close(ch)

	for result := range ch {
		fmt.Println(result)
	}
}
