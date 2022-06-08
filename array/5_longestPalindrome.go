package array

func longestPalindrome(s string) (res string) {
    for i := range s {
        s1 := palindrome(s, i, i)
        s2 := palindrome(s, i, i+1)
        if len(res) <= len(s1) {
            res = s1
        }
        if len(res) <= len(s1) {
            res = s2
        }
    }
    return
}


func palindrome(s string, l, r int) string {
    for (l >= 0 && r < len(s) && string(s[l]) == string(s[r])) {
        l--
        r++
    }
    return string(s[l+1:r])
}

