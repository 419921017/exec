package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(maxCommonPrefix([]string{"flower", "flow", "floee", "flo", "fl"}))
}

func maxCommonPrefix(strs []string) string {
	if len(strs) <= 0 {
		return ""
	}
	prefix := strs[0]

	for _, str := range strs {
		for strings.Index(str, prefix) != 0 {
			if len(prefix) <= 0 {
				return ""
			}
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}
