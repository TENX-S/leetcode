package binarySearch

func fairCandySwap(aliceSizes, bobSizes []int) (ans []int) {
	quickSort(aliceSizes)
	quickSort(bobSizes)

	aliceSum := arrSum(aliceSizes)
	bobSum := arrSum(bobSizes)
	avr := (aliceSum + bobSum) / 2
	for _, a := range aliceSizes {
		target := a + avr - aliceSum
		res := bSearch(bobSizes, target)
		if res != -1 {
			return []int{a, bobSizes[res]}
		}
	}

	return []int{}
}

func arrSum(arr []int) (res int) {
	for _, i := range arr {
		res += i
	}
	return
}

func bSearch(arr []int, target int) int {
	left, right := 0, len(arr)

	for left < right {
		mid := left + (right-left)/2
		key := arr[mid]
		if key == target {
			return mid
		} else if key < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return -1
}
