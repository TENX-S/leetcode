package Array

import "fmt"

func rotate(matrix [][]int) {
	d := len(matrix)
	for i := 0; i < d; i++ {
		for j := 0; j < d/2; j++ {
			a := matrix[i][j]
			b := matrix[j][i]
			matrix[i][j] = b
			matrix[j][i] = a
		}
	}

	fmt.Println(matrix)
	for r := 0; r < d; r++ {
		reverse(matrix[r])
	}
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i != j; {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}
