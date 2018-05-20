package main

import (
	"fmt"
)

// 给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。

// 示例 1:

// 输入: [1,2,3,4,5,6,7] 和 k = 3
// 输出: [5,6,7,1,2,3,4]
// 解释:
// 向右旋转 1 步: [7,1,2,3,4,5,6]
// 向右旋转 2 步: [6,7,1,2,3,4,5]
// 向右旋转 3 步: [5,6,7,1,2,3,4]
//解决思路，每反转一次，依次向后移动一位，末尾变为行首，依次循环而已！

func rotate(nums []int, k int) {

	for j:=1;j<=k;j++{

		size := len(nums)
		last := nums[size-1]
		for i := size - 1; i > 0; i-- {
			nums[i] = nums[i-1]
	
		}
		nums[0] = last
		
		fmt.Println(nums)

	}
	fmt.Println("last nums is", nums)

}

func main() {
	var nums = []int{1, 2, 3, 4, 5, 6,7,8,9}
	rotate(nums, 5)
}
