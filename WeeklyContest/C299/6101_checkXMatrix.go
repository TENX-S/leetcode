package C299

func checkXMatrix(grid [][]int) bool {
	n := len(grid[0])
	for i := 0; i < n; i++ {
		if grid[i][i] == 0 || grid[i][n-1-i] == 0 {
			return false
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j && i+j != n-1 {
				if grid[i][j] != 0 {
					return false
				}
			}
		}
	}
	return true
}
