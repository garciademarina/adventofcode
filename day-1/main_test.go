package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	expected := 7

	result := SolvePart1("/input-test.txt")

	if result != expected {
		t.Errorf("Result not expected %d != %d", result, expected)
	}

}

func TestPart2(t *testing.T) {
	expected := 5

	result := SolvePart2("/input-test.txt")

	if result != expected {
		t.Errorf("Result not expected %d != %d", result, expected)
	}

}
