package day1_test

import (
	"aoc/day1"
	"testing"
)

var input = []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

func TestSolvePart1(t *testing.T) {
	t.Parallel()

	expect := 7
	got := day1.SolvePart1(input)

	if expect != got {
		t.Fatalf(`expect: %v, got: %v`, expect, got)
	}
}

func TestSolvePart2(t *testing.T) {
	t.Parallel()

	expect := 5
	got := day1.SolvePart2(input)

	if expect != got {
		t.Fatalf(`expect: %v, got: %v`, expect, got)
	}
}
