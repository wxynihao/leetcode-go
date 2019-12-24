package main

import "fmt"

func main42() {
	arr1 := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	arr2 := []int{2, 0, 2}
	fmt.Println(trap(arr1))
	fmt.Println(trap(arr2))
}

func trap(height []int) int {
	length := len(height)
	if length < 2 {
		return 0
	}
	minWallArr := make([]int, length-2)

	max := height[0]
	for i := 0; i < length-2; i++ {
		if height[i] > max {
			max = height[i]
		}
		minWallArr[i] = max
	}

	max = height[length-1]
	for i := length - 1; i > 1; i-- {
		if height[i] > max {
			max = height[i]
		}
		if max < minWallArr[i-2] {
			minWallArr[i-2] = max
		}
	}

	sum := 0
	for i := 1; i < length-1; i++ {
		h := minWallArr[i-1] - height[i]
		if h > 0 {
			sum += h
		}
	}
	return sum
}
