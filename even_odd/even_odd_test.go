package main

import "testing"

func TestIsEven(t *testing.T) {
	e := []int{2, 4}
	o := []int{1, 3}
	for _, num := range e {
		if !isEven(num) {
			t.Errorf("Expected %v to be considered even, but failed.", num)
		}
	}

	for _, num := range o {
		if isEven(num) {
			t.Errorf("Expected %v to be considered odd, but failed.", num)
		}
	}
}
