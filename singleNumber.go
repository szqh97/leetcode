package main

import "fmt"

func singleNumber(nums []int) int {

	temp := 0
	for _, v := range nums {
		temp = temp ^ v
	}
	return temp
}
func main() {
	fmt.Println("vim-go")
}
