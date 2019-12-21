package main

import "fmt"

func main189() {

	f := func(nums []int, k int) {
		fmt.Printf("%d %d\n", nums, k)
		rotate(nums, k)
		fmt.Printf("%d\n\n", nums)
	}

	f([]int{0, 1, 2, 3, 4, 5, 6}, 2)
	f([]int{0, 1, 2, 3, 4, 5, 6}, 3)
	f([]int{0, 1, 2, 3, 4, 5, 6}, 4)

}

func rotate(nums []int, k int) {
	length := len(nums)
	k = k % length
	c := 0
	for i := 0; c < length; i++ {
		pos := i
		prev := nums[i]
		t := 0
		for i != pos || t == 0 {
			next := (pos + k) % length
			tmp := nums[next]
			nums[next] = prev
			prev = tmp
			pos = next
			c++
			t++
		}
	}
}
