package pascal

func Triangle(n int) [][]int {
	triangle := make([][]int, n)

	for i := 0; i < n; i++ {
		// The size of the inner array
		// is the row number.
		row := make([]int, i+1)
		for j := 0; j <= i; j++ {
			// The first and last elements in the row
			// are always 1
			if j == 0 || j == i {
				row[j] = 1
				continue
			}

			// The number in current position is
			// calculated from previous row
			row[j] = triangle[i-1][j-1] + triangle[i-1][j]
		}

		triangle[i] = row
	}

	return triangle
}
