package SfTech

func findMaxCI(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var stub int
	var p int
	ci := 1
	for p+1 < len(nums) {
		if nums[p+1] > nums[p] {
			stub = p
			for p+1 < len(nums) {
				if nums[p+1] > nums[p] {
					if p == len(nums)-2 {
						ci = max(ci, p-stub+2)
					}
					p++
				} else {
					ci = max(ci, p-stub+1)
					break
				}
			}
		}
		p++
	}

	return ci
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
