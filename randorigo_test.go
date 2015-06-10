package main

import (
	"testing"
	"time"
)

func getIdentity(i int) int { return i }
func getInverse(i int) int  { return -i }
func getInverseSlow(i int) int {
	time.Sleep (100*time.Millisecond)
	return -i
}

func TestRandomNumber(t *testing.T) {
	number := random()
	if number < 0 {
		t.Error("Random number should be greater than 0")
	}

	if number > 99999 {
		t.Error("Random number should be less than 99999")
	}
}

func TestRandomNumberSequence(t *testing.T) {
	sequence := randomSequence(100)
	if len(sequence) != 100 {
		t.Error("sequence length should be 100")
	}

	for _, n := range sequence {
		if n < 0 || n > 999999 {
			t.Error("number is not in range")
		}
	}
}

func TestSameInterval(t *testing.T) {
	min := minFinder(1, 1, getIdentity)
	if min != 1 {
		t.Error("min should be 1")
	}
}

func TestIncreasingInterval(t *testing.T) {
	min := minFinder(10, 20, getIdentity)
	if min != 10 {
		t.Error("min should be 10")
	}
}

func TestMaximumLessThanMinimum(t *testing.T) {
	min := minFinder(20, 10, getIdentity)
	if min != 10 {
		t.Error("min should be 10")
	}
}

func TestReturnsMinimumWithInverse(t *testing.T) {
	min := minFinder(20, 10, getInverse)
	if min != -20 {
		t.Error("min should be -20")
	}
}

func TestSquares(t *testing.T) {
	min := minFinder(-1, 1, func(i int) int { return i * i })
	if min != 0 {
		t.Error("min should be 0")
	}
}

func Test200msBasic(t *testing.T) {
	startTime := time.Now()
  minFinder(1, 10, getInverseSlow)
  if time.Since(startTime) * time.Millisecond > 200 {
		t.Error("time greater than 200ms")
	}
}
