package binarySearch

type Row struct {
	idx, num int
}

func (rec Row) Less(other Row) bool {
	if rec.num < other.num {
		return true
	}
	if rec.num == other.num && rec.idx < other.idx {
		return true
	}
	return false
}

func (rec Row) Equal(other Row) bool {
	if rec.idx == other.idx && rec.num == other.num {
		return true
	}
	return false
}

func (rec Row) Greater(other Row) bool {
	return !rec.Less(other) && !rec.Equal(other)
}

func kWeakestRows(mat [][]int, k int) []int {
	recs := make([]Row, len(mat))
	for i, arr := range mat {
		recs[i] = Row{idx: i}
		cnt := bSearchWithRightBound(arr, 1)
		if cnt != -1 {
			recs[i].num = cnt + 1
		}
	}
	rowQSort(recs)
	ans := make([]int, k)
	for i := 0; i < k; i++ {
		ans[i] = recs[i].idx
	}
	return ans
}

func bSearchWithRightBound(arr []int, target int) int {
	left, right := 0, len(arr)

	for left < right {
		mid := left + (right-left)/2
		key := arr[mid]
		if key == target {
			left = mid + 1
		} else if key > target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if left == 0 || arr[left-1] != target {
		return -1
	}

	return left - 1
}

func rowQSort(recs []Row) {
	rowSort(recs, 0, len(recs)-1)
}

func rowSort(recs []Row, start, end int) {
	if start < end {
		i, j := start, end
		key := recs[(start+end)/2]
		for i <= j {
			for recs[i].Less(key) {
				i++
			}
			for recs[j].Greater(key) {
				j--
			}
			if i <= j {
				recs[i], recs[j] = recs[j], recs[i]
				i++
				j--
			}
		}
		if start < j {
			rowSort(recs, start, j)
		}
		if end > i {
			rowSort(recs, i, end)
		}
	}
}
