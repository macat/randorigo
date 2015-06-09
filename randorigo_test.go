package main

import "testing"

func TestRandomNumber(t *testing.T) {
	number := random()
	if number < 0 {
		t.Error("random number should be greater than 0")
	}
	if number > 999999 {
		t.Error("random number should be smaller than 999999")
	}
}

func TestRandomSequence(t *testing.T) {
	sequence := randomSequence(100)

	if len(sequence) != 100 {
		t.Error("sequence length is not corrent")
	}

	sum := 0

	for _, number := range sequence {
		sum += number
		if number < 0 {
			t.Error("random number should be greater than 0")
		}
		if number > 999999 {
			t.Error("random number should be smaller than 999999")
		}
	}

	avg := sum / 100
	if avg == sequence[10] && avg == sequence[30] {
		t.Error("sequence is not random")
	}
}
