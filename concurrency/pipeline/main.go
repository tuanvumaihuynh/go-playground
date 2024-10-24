package main

import "fmt"

func sliceToChan(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	// input
	nums := []int{2, 3, 4, 9, 1, 3, 4}

	// stage 1
	dataCh := sliceToChan(nums)
	// stage 2
	finalCh := sq(dataCh)
	// stage 3
	for n := range finalCh {
		fmt.Println(n)
	}
}
