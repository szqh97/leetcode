package main

import "fmt"

func removeDuplicates(nums []int) int {
	mp := make(map[int]int)
	for _, v := range nums {
		if _, ok := mp[v]; !ok {
			mp[v] = len(mp)
			nums[mp[v]] = v
		}

	}
	return len(mp)
}
func main() {
	//nums := []int{1, 1, 2}
	//nums := []int{1, 2, 3, 4, 4, 4, 4, 5}
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)

}
