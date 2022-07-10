package Sort

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, start, end int) {
	if start < end {
		i, j := start, end
		key := arr[(start+end)/2]
		for i <= j {
			for arr[i] < key {
				i++
			}
			for arr[j] > key {
				j--
			}
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}

		if start < j {
			quickSort(arr, start, j)
		}
		if end > i {
			quickSort(arr, i, end)
		}
	}
}

func quickSort2(q []int, l, r int) {
	if l >= r {
		return
	}
	x := q[(l+r)>>1]
	i, j := l-1, r+1
	for i < j {
		for {
			i++
			if q[i] >= x {
				break
			}
		}
		for {
			j--
			if q[j] <= x {
				break
			}
		}
		if i < j {
			q[i], q[j] = q[j], q[i]
		}
	}
	quickSort2(q, l, j)
	quickSort2(q, j+1, r)
}

// func main() {
//     var n int
//     fmt.Scanf("%d", &n)
//     q := make([]int, n)
//     for i := 0; i < n; i ++ {
//         fmt.Scanf("%d", &q[i])
//     }
//     quickSort(q, 0, n-1)
//     for i := 0; i < n; i ++ {
//         fmt.Printf("%d ",q[i] )
//     }
// }
//
