package main

import (
	"fmt"
	"strconv"
	"strings"
)

// CountAndSay returns the n-th term of the look-and-say sequence: each term
// reads the previous one aloud as run-lengths ("21" -> "one 2, one 1" -> "1211").
func CountAndSay(n int) string {
	if n < 1 {
		return ""
	}
	s := "1"
	for i := 1; i < n; i++ {
		var b strings.Builder
		for j := 0; j < len(s); {
			k := j
			for k < len(s) && s[k] == s[j] {
				k++
			}
			b.WriteString(strconv.Itoa(k - j))
			b.WriteByte(s[j])
			j = k
		}
		s = b.String()
	}
	return s
}

func main() {
	for n := 1; n <= 6; n++ {
		fmt.Println(CountAndSay(n))
	}
}
