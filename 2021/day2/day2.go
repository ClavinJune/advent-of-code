package day2

import (
	"fmt"
)

func SolvePart1(arr []string) int {
	var h, d int
	for _, cmd := range arr {
		var pos string
		var unit int

		fmt.Sscanf(cmd, "%s %d", &pos, &unit)

		switch pos {
		case "forward":
			h += unit
		case "down":
			d += unit
		case "up":
			d -= unit
		}
	}
	return h * d
}

func SolvePart2(arr []string) int {
	var h, d, a int
	for _, cmd := range arr {
		var pos string
		var unit int

		fmt.Sscanf(cmd, "%s %d", &pos, &unit)

		switch pos {
		case "forward":
			h += unit
			d += (a * unit)
		case "down":
			a += unit
		case "up":
			a -= unit
		}
	}
	return h * d
}
