package main

import "fmt"

func main() {
	//strs := []string{"abc", "abcd", "abck", "a"}
	strs := []string{"1", ""}

	fmt.Println(longestCommonPrefix(strs))

}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 || strs[0] == "" {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	result := ""
	c := strs[0][0]
	for i := 0; ; i++ {
		if len(strs[0]) > i {
			c = strs[0][i]
		} else {
			break
		}
		for k, _ := range strs {
			str := strs[k]
			if len(str) <= i {
				return result
			}

			if c == str[i] && k == len(strs)-1 {
				result += string(c)
			}
			if c != str[i] {
				return result
			}
		}

	}
	return string(result)

}
