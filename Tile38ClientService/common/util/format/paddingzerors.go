package format

import "fmt"

func PadingZerors19(num int64) string {
	s := fmt.Sprintf("%019d", num)
	return s
}

func PadingZerors10(num int64) string {
	s := fmt.Sprintf("%010d", num)
	return s
}
