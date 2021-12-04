package day3_test

import (
	"aoc/day3"
	"testing"
)

var input = []string{
	"00100",
	"11110",
	"10110",
	"10111",
	"10101",
	"01111",
	"00111",
	"11100",
	"10000",
	"11001",
	"00010",
	"01010",
}

func TestSolvePart1(t *testing.T) {
	t.Parallel()

	expect := 198
	got := day3.SolvePart1(input)

	if expect != got {
		t.Fatalf(`expect: %v, got: %v`, expect, got)
	}
}

func TestSolvePart2(t *testing.T) {
	t.Parallel()

	expect := 230
	got := day3.SolvePart2(input)

	if expect != got {
		t.Fatalf(`expect: %v, got: %v`, expect, got)
	}
}
