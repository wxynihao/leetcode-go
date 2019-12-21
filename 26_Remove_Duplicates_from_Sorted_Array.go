package main

import "fmt"

func main() {

	f := func(nums []int) {
		fmt.Println(nums)
		fmt.Printf("%d\nlength:%d\n\n",
			nums, removeDuplicates(nums))
	}

	f([]int{1, 1, 2})
	f([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
	f([]int{0})
	f([]int{})
	f([]int{1, 2, 3, 4})
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
