package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	a := 1
	b := 2
	expected := 3
	result := sum(a, b)
	if result != expected {
		t.Errorf("Sum(%d, %d) = %d; want %d", a, b, result, expected)
	}
}

func TestSumb(t *testing.T) {
	a := 5
	b := 2
	expected := 3
	result := sub(a, b)
	if result != expected {
		t.Errorf("Sum(%d, %d) = %d; want %d", a, b, result, expected)
	}
}

func TestTimes(t *testing.T) {
	a := 2
	b := 2
	expected := 4
	result := times(a, b)
	if result != expected {
		t.Errorf("Times(%d, %d) = %d; want %d", a, b, result, expected)
	}
}

func TestSumX(t *testing.T) {
	a := 2
	b := 2
	expected := 6
	result := sumX(a, b)
	if result != expected {
		t.Errorf("SumX(%d, %d) = %d; want %d", a, b, result, expected)
	}
}

func TestMain(t *testing.T) {
	main()
}
