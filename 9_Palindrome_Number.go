package main

import "fmt"

func main9() {
	f := func(x int) {
		fmt.Printf("%d %g\n", x, isPalindrome(x))
	}

	f(-1)
	f(0)
	f(-0)
	f(1)
	f(123)
	f(121)
	f(1221)

}

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	y := 0
	for x > y {
		y = y*10 + x%10
		x /= 10
	}
	return x == y || y/10 == x
}
