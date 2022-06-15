package BinarySearch

func findMin(nums []int) int {
	peek := findPeek(nums)
	if peek == len(nums)-1 {
		return nums[0]
	} else {
		return nums[peek+1]
	}
}
