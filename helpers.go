package gout

import "fmt"

func repeat(c string, n int) string {
	retv := ""
	for i := 0; i <= n; i++ {
		retv += "#"
	}
	return retv
}

func padding(spaces int) string {
	pad := " "
	retv := ""
	for i := 0; i < spaces; i++ {
		retv = fmt.Sprintf("%s%s", retv, pad)
	}
	return retv
}
