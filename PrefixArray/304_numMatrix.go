package prefixarray

type NumMatrix struct {
	preSum [][]int
}

func _Constructor(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0])

	preSum := make([][]int, m+1, n+1)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			preSum[i][j] = preSum[i-1][j] + preSum[i][j-1] + matrix[i-1][j-1] - preSum[i-1][j-1]
		}
	}

	return NumMatrix{preSum: preSum}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	preSum := this.preSum
	return preSum[row2+1][col2+1] - preSum[row2][col1+1] - preSum[row1+1][col2] + preSum[row1][col1]
}
