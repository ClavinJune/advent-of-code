package day1

func SolvePart1(arr []int) int {
	n := len(arr)

	var c int
	for i := 1; i < n; i++ {
		if arr[i] > arr[i-1] {
			c++
		}
	}

	return c
}

func SolvePart2(arr []int) int {
	n := len(arr)

	var c int
	for i := 1; i < n-2; i++ {
		window := arr[i] + arr[i+1]
		prev := arr[i-1] + window
		curr := window + arr[i+2]
		if curr > prev {
			c++
		}
	}

	return c
}
