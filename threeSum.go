package main

import "fmt"

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

type MinMax struct {
	Min int
	Max int
}

// a + b + c = 0
// c is caculated
func threeSum(nums []int) [][]int {
	result := [][]int{}
	gmp := map[MinMax]int{}
	for i := 0; i < len(nums); i++ {
		a := nums[i]
	}
	for k, v := range gmp {
		result = append(result, []int{k.Min, k.Max, v})
	}

	return result

}
func main() {
	//nums := []int{-1, 0, 1, 0, 2, -1, -3, -5, 4}
	//nums := []int{0, 0, 0, 0}
	//nums := []int{-1, 0, 1, 2, -1, -4}
	//nums := []int{2, -1, -1, 0}
	//fmt.Println(twoSum(nums, -1))
	nums := []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}
	fmt.Println(threeSum(nums))

}
