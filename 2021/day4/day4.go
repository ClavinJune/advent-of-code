package day4

type boardMap struct {
	n, x, y int
}

func bingo(arr []int, m map[int][]*boardMap, r map[int][]int, s map[int]int) int {
	for _, a := range arr {
		for _, v := range m[a] {
			s[v.n] -= a
			col := r[v.n][v.x] - a
			if col == 0 {
				r[v.n][v.x] = col
				return a * s[v.n]
			} else {
				r[v.n][v.x] = col
			}

			row := r[v.n][v.y+5] - a
			if col == 0 {
				r[v.n][v.y+5] = row
				return a * s[v.n]
			} else {
				r[v.n][v.y+5] = row
			}
		}
	}
	return -1
}

func SolvePart1(arr []int, boards [][5][5]int) int {
	m := make(map[int][]*boardMap)
	nb := len(boards)

	for _, a := range arr {
		m[a] = make([]*boardMap, 0, nb)
	}

	r := make(map[int][]int)
	s := make(map[int]int)
	for i := range boards {
		r[i] = make([]int, 10)
		for j := range boards[i] {
			for k, num := range boards[i][j] {
				if _, ok := m[num]; !ok {
					continue
				}

				m[num] = append(m[num], &boardMap{i, j, k})
				r[i][j] += num
				r[i][k+5] += num
				s[i] += num
			}
		}
	}
	return bingo(arr, m, r, s)
}
