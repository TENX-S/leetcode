package prefixarray

type NumMatrix struct {
	nums [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0])
	nums := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		nums[i] = make([]int, n+1)
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			nums[i][j] = nums[i-1][j] + nums[i][j-1] - nums[i-1][j-1] + matrix[i][j]
		}
	}
	return NumMatrix{nums: nums}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	nums := this.nums
	return nums[row2][col2] - nums[row1-1][col2] - nums[row2][col1-1] + nums[row1-1][col1-1]
}
