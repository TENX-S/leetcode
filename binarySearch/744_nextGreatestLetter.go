package binarySearch

func nextGreatestLetter(letters []byte, target byte) byte {
    left, right := 0, len(letters)

    for left < right {
        mid := left + (right-left)/2
        letter := letters[mid]
        if letter > target {
            right = mid
        } else if letter == target {
            left = mid + 1
        } else {
            left = mid + 1
        }
    }

    if left == len(letters) {
        return letters[0]
    }

    return letters[left]
}
