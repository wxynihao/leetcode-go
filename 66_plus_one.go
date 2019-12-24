package main

import "fmt"

func main66() {
	f := func(digits []int) {
		fmt.Println(plusOne(digits))
	}
	f([]int{1, 2, 9})
	f([]int{9})
}

func plusOne(digits []int) []int {
	length := len(digits)
	flag := 1
	for i := length - 1; i >= 0; i-- {
		if flag != 1 {
			break
		}
		digits[i] += 1
		if digits[i] != 10 {
			flag = 0
			break
		}
		digits[i] = 0
	}

	if flag == 1 {
		digits = append(digits, 0)
		copy(digits[1:], digits)
		digits[0] = 1
	}

	return digits
}
