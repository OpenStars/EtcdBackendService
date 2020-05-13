package format

import "strconv"

func S2I(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}
