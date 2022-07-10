package C300

import "strings"

func decodeMessage(key string, message string) string {
	m := make(map[string]string)
	key = strings.ReplaceAll(key, " ", "")
	for i, j := 0, 0; i < len(key); i++ {
		k := string(key[i])
		if j > 25 {
			break
		}
		if _, exist := m[k]; !exist {
			n := string(rune(97 + j))
			m[string(key[i])] = n
			j++
		}
	}
	translated := ""
	for i, b := range []byte(message) {
		if b != byte(' ') {
			translated += m[string(message[i])]
		} else {
			translated += " "
		}
	}
	return string(translated)
}
