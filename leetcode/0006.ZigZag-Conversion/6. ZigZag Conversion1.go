package leetcode

func convert1(s string, numRows int) string {
	if numRows == 1 || numRows > len(s) {
		return s
	}
	slice, down, up := make([][]byte, numRows), 0, numRows-2
	for i := 0; i < len(s); {
		if down < numRows {
			slice[down] = append(slice[down], s[i])
			down++
			i++
		} else if up > 0 {
			slice[up] = append(slice[up], s[i])
			up--
			i++
		} else {
			up = numRows - 2
			down = 0
		}
	}
	solution := make([]byte, 0, len(s))
	for _, row := range slice {
		solution = append(solution, row...)
	}
	return string(solution)
}
