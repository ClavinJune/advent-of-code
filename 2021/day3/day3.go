package day3

import (
	"strconv"
)

func SolvePart1(arr []string) int {
	c := make([][2]int, len(arr[0]))

	for i := range arr {
		for j := range arr[i] {
			c[j][arr[i][j]-48]++
		}
	}

	var g, e string
	for i := range c {
		if c[i][0] > c[i][1] {
			g += "0"
			e += "1"
		} else {
			g += "1"
			e += "0"
		}
	}

	gi, _ := strconv.ParseInt(g, 2, 64)
	ei, _ := strconv.ParseInt(e, 2, 64)

	return int(gi * ei)
}

func findOx(m map[string]struct{}, i int) string {
	var c [2]int
	for k := range m {
		c[k[i]-48]++
	}

	b := "1"
	if c[0] > c[1] {
		b = "0"
	}

	for k := range m {
		if k[i] != b[0] {
			delete(m, k)
		}
	}
	return b
}

func findCo(m map[string]struct{}, i int) string {
	var c [2]int
	for k := range m {
		c[k[i]-48]++
	}

	b := "0"
	if c[0] > c[1] {
		b = "1"
	}

	for k := range m {
		if k[i] != b[0] {
			delete(m, k)
		}
	}

	return b
}

func SolvePart2(arr []string) int {
	mox := make(map[string]struct{})
	mco := make(map[string]struct{})

	for _, a := range arr {
		mox[a] = struct{}{}
		mco[a] = struct{}{}
	}

	n := len(arr[0])
	for i := 0; i < n; i++ {
		if len(mox) > 1 {
			findOx(mox, i)
		}

		if len(mco) > 1 {
			findCo(mco, i)
		}
	}

	var ox, co string
	for k := range mox {
		ox = k
		break
	}
	for k := range mco {
		co = k
		break
	}

	oxi, _ := strconv.ParseInt(ox, 2, 64)
	coi, _ := strconv.ParseInt(co, 2, 64)

	return int(oxi * coi)

	return -1
}
