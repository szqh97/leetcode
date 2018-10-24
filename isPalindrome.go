package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	origin := x
	result := 0
	for origin > result {
		result = result*10 + x%10
		x = x / 10
	}
	return origin == result || origin == result/10
}

func main() {
	fmt.Println(isPalindrome(123))
	fmt.Println(isPalindrome(1230))
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))

}
