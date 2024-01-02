package main

import "testing"

func Test_If_Sum_Is_Correct(t *testing.T) {
	result := Sum(10, 5, 8)
	if result != 23 {
		t.Errorf("Sum was incorrect, result: %d, want: %d.", result, 23)
	}
}

func Test_If_Multiply_Is_Correct(t *testing.T) {
	result := Multiply(10, 5, 8)
	if result != 400 {
		t.Errorf("Multiply was incorrect, result: %d, want: %d.", result, 400)
	}
}

func Test_If_Subtract_Is_Correct(t *testing.T) {
	result := Subtract(10, 5, 8)
	if result != -3 {
		t.Errorf("Subtract was incorrect, result: %d, want: %d.", result, -3)
	}
}

func Test_If_Divide_Is_Correct(t *testing.T) {
	result := Divide(10, 5, 8)
	if result != 0.25 {
		t.Errorf("Divide was incorrect, result: %f, want: %f.", result, 0.25)
	}
}

func Test_If_Divide_Is_Incorrect(t *testing.T) {
	result := Divide(10, 5, 0)
	if result != 0 {
		t.Errorf("Divide was incorrect, result: %v, want: %v.", result, 0)
	}
}