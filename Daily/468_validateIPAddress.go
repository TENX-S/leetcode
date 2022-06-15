package Daily

import (
	"strconv"
	"strings"
)

func validIPAddress(queryIP string) string {
	if strings.Contains(queryIP, ".") {
		seg := strings.Split(queryIP, ".")
		if len(seg) != 4 {
			return "Neither"
		} else {
			for _, s := range seg {
				if strings.HasPrefix(s, "0") && len(s) > 1 {
					return "Neither"
				}
				if r, err := strconv.Atoi(s); err != nil {
					return "Neither"
				} else {
					if r < 0 || r > 255 {
						return "Neither"
					}
				}
			}
			return "IPV4"
		}
	} else if strings.Contains(queryIP, ":") {
		seg := strings.Split(queryIP, ":")
		if len(seg) != 8 {
			return "Neither"
		} else {
			for _, r := range seg {
				sLen := len(r)
				r = strings.ToLower(r)
				if sLen < 1 || sLen > 4 {
					return "Neither"
				}
				if _, err := strconv.ParseUint(r, 16, 16); err != nil {
					return "Neither"
				}
			}
			return "IPV6"
		}
	}
	return "Neither"
}
