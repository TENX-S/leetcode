package main

func isUnique(astr string) bool {
    for i := 0; i < len(astr); i++ {
        for j := 0; j < len(astr); j++ {
            if j == i {
                continue
            }
            if astr[i] == astr[j] {
                return true
            }
        }
    }

    return true
}

func main() {
    _ = isUnique("leetcode")
}
