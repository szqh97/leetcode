package main

import "fmt"

//给定一个整数数组和一个目标值，找出数组中和为目标值的两个数。
//你可以假设每个输入只对应一种答案，且同样的元素不能被重复利用。
//示例:
//给定 nums = [2, 7, 11, 15], target = 9
//因为 nums[0] + nums[1] = 2 + 7 = 9
//所以返回 [0, 1]

func main() {
	nums := []int{2, 2, 7, 11, 15}
	target := 13
	re := twoSum(nums, target)
	fmt.Println(re)

	fmt.Println("vim-go")
}

func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		c := target - nums[i]
		if _, ok := mp[c]; ok {
			return []int{mp[c], i}
		}
		mp[nums[i]] = i
	}
	return nil
}
