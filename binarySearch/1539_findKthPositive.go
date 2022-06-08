package binarySearch

type Tag struct {
	val     int
	existed bool
}

func findKthPositive(arr []int, k int) int {
	seq := make([]Tag, arr[len(arr)-1]+k)
	for i := range seq {
		seq[i] = Tag{
			val:     i + 1,
			existed: false,
		}
	}

	for _, a := range arr {
		seq[a-1].existed = true
	}

	var p int

	for _, e := range seq {
		if !e.existed {
			p++
		}
		if p == k {
			return e.val
		}
	}

	return -1
}
