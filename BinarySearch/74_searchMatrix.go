package BinarySearch

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	lC, rC := 0, m
	for lC < rC {
		mC := lC + (rC-lC)/2
		key := matrix[mC][n-1]
		if key == target {
			return true
		} else if key < target {
			lC = mC + 1
		} else if key > target {
			rC = mC
		}
	}
	if lC >= m {
		return false
	}

	row := lC
	l, r := 0, n
	for l < r {
		mid := l + (r-l)/2
		key := matrix[row][mid]
		if key == target {
			return true
		} else if key < target {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return false
}
