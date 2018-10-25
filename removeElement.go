package main

import "fmt"

func removeElement(nums []int, val int) int {

	i := 0
	j := len(nums) - 1
	for ; i <= j; i++ {
		for nums[j] == val {
			if i == j {
				return i
			}
			j = j - 1
		}
		if nums[i] == val {
			nums[i], nums[j] = nums[j], nums[i]
			j = j - 1
		}
	}
	return i
}
func main() {
	nums := []int{2}
	fmt.Println(removeElement(nums, 4))
	fmt.Println(nums)

}
