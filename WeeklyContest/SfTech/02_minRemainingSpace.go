package SfTech

import "sort"

func minRemainingSpace(N []int, V int) int {
	sort.Ints(N)
	Vp := V
	var i, j int
	for i = 0; i < len(N) && V-N[i] > N[len(N)-1]; i++ {
		V -= N[i]
	}
	if i == len(N)-1 {
		return Vp - arrSum(N)
	}
	if V < N[0] {
		return Vp
	}

	ans := Vp - arrSum(N[:i+1])
	for j = i + 1; j < len(N) && V >= N[j]; j++ {
	}

	if j < len(N) && N[j] <= ans {
		ans -= N[j-1]
	}

	if j == len(N) && N[j-1] <= ans {
		ans -= N[j-1]
	}
	return ans
}

func arrSum(arr []int) (res int) {
	for _, v := range arr {
		res += v
	}
	return
}
