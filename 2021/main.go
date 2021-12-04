package main

import (
	"aoc/day1"
	"aoc/day2"
	"aoc/day3"
	"fmt"
)

func solveDay1() {
	fmt.Println("day 1 part 1", day1.SolvePart1(day1.Input))
	fmt.Println("day 1 part 1", day1.SolvePart2(day1.Input))
}

func solveDay2() {
	fmt.Println("day 2 part 1", day2.SolvePart1(day2.Input))
	fmt.Println("day 2 part 2", day2.SolvePart2(day2.Input))
}

func solveDay3() {
	fmt.Println("day 3 part 1", day3.SolvePart1(day3.Input))
	fmt.Println("day 3 part 2", day3.SolvePart2(day3.Input))
}

func main() {
	solveDay1()
	solveDay2()
	solveDay3()
}
