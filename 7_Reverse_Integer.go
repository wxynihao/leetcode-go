package main

import "fmt"

func main7() {
	f := func(x int) {
		fmt.Printf("%d %d\n", x, reverse(x))
	}

	f(123)
	f(-123)
	f(120)
	f(2147483647)
	f(-2147483648)
	f(1534236469)
}

func reverse(x int) int {
	y := int32(x)
	var r int32
	for r = 0; y != 0; y /= 10 {
		tmp := y % 10
		t := 10 * r
		if t/10 != r {
			return 0
		}
		r = t + tmp
	}
	return int(r)
}
