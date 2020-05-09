package main

import "strings"

func lengthOfLastWord(s string) int {
	s = strings.Trim(s, " ")
	if len(s) == 0 {
		return 0
	}
	strArr := strings.Split(s, " ")
	lenOfArr := len(strArr)
	if lenOfArr == 0 {
		return 0
	}
	return len(strArr[lenOfArr-1])
}
