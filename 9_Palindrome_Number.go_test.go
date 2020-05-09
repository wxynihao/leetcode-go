package main

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	f := func(x int) {
		fmt.Printf("%d %t\n", x, isPalindrome(x))
	}

	f(-1)
	f(0)
	f(-0)
	f(1)
	f(123)
	f(121)
	f(1221)
}
