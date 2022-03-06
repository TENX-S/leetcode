package array

// https://leetcode-cn.com/problems/sorted-merge-lcci/solution/mian-shi-ti-1001-he-bing-pai-xu-de-shu-zu-by-leetc/

func merge(A []int, m int, B []int, n int) {
    sorted := make([]int, m+n)
    pa, pb := 0, 0

    for {
        if pa == m {
            sorted = append(sorted, B[pb:]...)
            break
        }
        if pb == n {
            sorted = append(sorted, A[pa:]...)
            break
        }
        if A[pa] <= B[pb] {
            sorted = append(sorted, A[pa])
            pa++
        } else {
            sorted = append(sorted, B[pb])
            pb++
        }
    }

    copy(A, sorted)
}
