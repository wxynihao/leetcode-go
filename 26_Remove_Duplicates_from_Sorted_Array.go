package main

import "fmt"

func main() {
	nums := []int{1, 1, 2}
	fmt.Printf("length:%d\n", removeDuplicates(nums))
	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Printf("length:%d\n", removeDuplicates(nums))
	nums = []int{0}
	fmt.Printf("length:%d\n", removeDuplicates(nums))
	nums = []int{}
	fmt.Printf("length:%d\n", removeDuplicates(nums))
	nums = []int{1, 2, 3, 4}
	fmt.Printf("length:%d\n", removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	pos := 0
	tmp := nums[0]
	for i := 1; i < length; i++ {
		if nums[i] != tmp {
			pos++
			tmp = nums[i]
			if pos != i {
				nums[pos] = nums[i]
			}
		}
	}
	return pos + 1
}
