package C301

import "math"

func fillCups(amount []int) int {
	var i int
	for amount[0] != 0 && amount[1] != 0 && amount[2] != 0 {
		amount[maxArr(amount)]--
		amount[minArr(amount)]--
		i++
	}
	return i + amount[maxArr(amount)]
}

func maxArr(arr []int) int {
	max := math.MinInt
	var curr int
	for i, v := range arr {
		if v > max {
			max = v
			curr = i
		}
	}
	return curr
}

func minArr(arr []int) int {
	min := math.MaxInt
	var curr int
	for i, v := range arr {
		if v < min {
			min = v
			curr = i
		}
	}
	return curr
}
