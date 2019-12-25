package main

import "fmt"

func main283() {
	arr := []int{0, 1, 0, 3, 12}
	moveZeroes(arr)
	fmt.Println(arr)

}

func moveZeroes(nums []int) {
	length := len(nums)
	posOfJ := 1
	for i := 0; i < length; i++ {
		if nums[i] != 0 {
			continue
		}
		if i+1 > posOfJ {
			posOfJ = i + 1
		}
		for j := posOfJ; j < length; j++ {
			if nums[j] == 0 {
				continue
			}
			nums[i] = nums[j]
			nums[j] = 0
			posOfJ = j + 1
			break
		}
	}
}

func moveZeroes2(nums []int)  {
	pos := 0
	for i, tmp := range nums {
		if tmp != 0 {
			if i!= pos{
				nums[pos] = tmp
			}
			pos++
		}
	}
	for pos < len(nums) {
		nums[pos] = 0
		pos++
	}
}
