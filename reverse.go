package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.MinInt32)

	fmt.Println(reverse(113))
	fmt.Println(reverse(111000))
	fmt.Println(reverse(-111000))
	fmt.Println(reverse(-214748364))
	fmt.Println(reverse(-2147483412))
	fmt.Println(-2147483412 >= math.MinInt32)
	fmt.Println(-2143847412 >= math.MinInt32)

}

func reverse(x int) int {
	rem := 0
	var result int = 0
	for x != 0 {
		rem = x % 10
		if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && rem > 7) {
			return 0
		} else if result < math.MinInt32/10 || (result == math.MinInt32/10 && rem < -8) {
			return 0
		}
		result = result*10 + rem
		x = x / 10
	}
	return result

}
