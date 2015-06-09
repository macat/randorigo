package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("hello world")
}

func random() int {
	time.Sleep(100 * time.Millisecond)
	return rand.Intn(999999)
}

func randomSequence(length int) []int {
	sequence := []int{}
	randomChannel := make(chan int, 10)

	for i := 0; i < length; i++ {
		go func() {
			randomChannel <- random()
		}()
	}

	for i := 0; i < length; i++ {
		number := <-randomChannel
		sequence = append(sequence, number)
	}

	return sequence
}
