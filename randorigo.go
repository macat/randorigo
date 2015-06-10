package main

import (
	"math/rand"
	"time"
)

const MAX_VALUE int = 99999

func main() {
}

func random() int {
	time.Sleep(100 * time.Millisecond)
	return rand.Intn(MAX_VALUE)
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

func minFinder(begin int, end int, f func(int) int) int {
	min := f(begin)
	if begin > end {
		begin, end = end, begin
	}
	for i := begin; i < end; i++ {
		temp := f(i)
		if temp < min {
			min = temp
		}
	}

	return min
}
