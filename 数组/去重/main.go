package main

import (
	"fmt"
)

// 给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。

// 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。

// 示例 1:

// 给定数组 nums = [1,1,2],

// 函数应该返回新的长度 2, 并且原数组 nums 的前两个元素被修改为 1, 2。

// 你不需要考虑数组中超出新长度后面的元素。

//解决思路:首先数组是有序的，我们只需相邻的相互依次比较，元素相同则比较下一个，元素不同则添加到新的数组中
func removeDuplicates(nums []int) int {
	var newNums []int
	l := len(nums)
	for k, _ := range nums {
		if k+1 < l {
			if nums[k] == nums[k+1] {
				continue
			}
		}
		newNums = append(newNums, nums[k])
	}
	fmt.Println("newnums is ", newNums)
	fmt.Println("len is ", len(newNums))
	return len(newNums)
}

func main() {
	nums := []int{1, 1, 2, 2, 2, 2, 2, 2, 2, 3, 4, 5, 6, 7}
	removeDuplicates(nums)
}
