package binarySearch

func isPerfectSquare(num int) bool {
    if num <= 1 {
        return true
    }

    left, right := 2, num/2
    for left <= right {
        mid := left + (right-left)/2
        key := mid * mid
        if key == num {
            return true
        } else if key < num {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return false
}
