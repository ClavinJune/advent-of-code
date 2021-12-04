package day2_test

import (
	"aoc/day2"
	"testing"
)

var input = []string{
	"forward 5",
	"down 5",
	"forward 8",
	"up 3",
	"down 8",
	"forward 2",
}

func TestSolvePart1(t *testing.T) {
	t.Parallel()

	expect := 150
	got := day2.SolvePart1(input)

	if expect != got {
		t.Fatalf(`expect: %v, got: %v`, expect, got)
	}
}

func TestSolvePart2(t *testing.T) {
	t.Parallel()

	expect := 900
	got := day2.SolvePart2(input)

	if expect != got {
		t.Fatalf(`expect: %v, got: %v`, expect, got)
	}
}
