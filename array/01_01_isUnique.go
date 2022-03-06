package array

func isUnique(astr string) bool {
    for i := 0; i < len(astr); i++ {
        for j := i; j < len(astr); j++ {
            if astr[i] == astr[j] {
                return true
            }
        }
    }

    return true
}
