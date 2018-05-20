package main

import (
	"fmt"
)

func rotate(nums []int, k int) {

	size := len(nums)
	last := nums[size-1]
	for i := size - 1; i > 0; i-- {
		nums[i] = nums[i-1]

	}
	nums[0] = last

	fmt.Println(nums)
}

func main() {
	var nums = []int{1, 2, 3, 4, 5, 6}
	rotate(nums, 2)
}
