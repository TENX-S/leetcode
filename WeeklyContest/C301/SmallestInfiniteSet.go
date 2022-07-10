package C301

type SmallestInfiniteSet struct {
	min   int
	poped []int
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{min: 1}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	ret := this.min
	this.poped = append(this.poped, ret)
	this.min++
	for ElemIn(this.poped, this.min) {
		this.min++
	}
	return ret
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if num < this.min {
		this.min = num
	}
	if ElemIn(this.poped, num) {
		this.poped = removeElem(this.poped, num)
	}
}

func removeElem(arr []int, target int) []int {
	for i, v := range arr {
		if v == target {
			arr[i] = arr[len(arr)-1]
		}
	}
	arr = arr[:len(arr)-1]
	return arr
}

func ElemIn(arr []int, target int) bool {
	for _, v := range arr {
		if v == target {
			return true
		}
	}
	return false
}
