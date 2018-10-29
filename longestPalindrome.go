package main

import "fmt"

func palindrom(s string) bool {
	for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	pali_str := ""
	for i := 0; i < len(s); i++ {
		for j := 0; j <= i && i+j < len(s); j++ {
			if i-j >= 0 {
				tmps := s[i-j : i+j+1]
				if palindrom(tmps) && len(tmps) > len(pali_str) {
					pali_str = tmps
				}
			}
			if i-j-1 >= 0 {
				tmps2 := s[i-j-1 : i+j+1]

				if palindrom(tmps2) && len(tmps2) > len(pali_str) {
					pali_str = tmps2
				}
			}

		}
	}
	return pali_str

}
func main() {
	s := "cc1234"
	s = "azwdzwmwcqzgcobeeiphemqbjtxzwkhiqpbrprocbppbxrnsxnwgikiaqutwpftbiinlnpyqstkiqzbggcsdzzjbrkfmhgtnbujzszxsycmvipjtktpebaafycngqasbbhxaeawwmkjcziybxowkaibqnndcjbsoehtamhspnidjylyisiaewmypfyiqtwlmejkpzlieolfdjnxntonnzfgcqlcfpoxcwqctalwrgwhvqvtrpwemxhirpgizjffqgntsmvzldpjfijdncexbwtxnmbnoykxshkqbounzrewkpqjxocvaufnhunsmsazgibxedtopnccriwcfzeomsrrangufkjfzipkmwfbmkarnyyrgdsooosgqlkzvorrrsaveuoxjeajvbdpgxlcrtqomliphnlehgrzgwujogxteyulphhuhwyoyvcxqatfkboahfqhjgujcaapoyqtsdqfwnijlkknuralezqmcryvkankszmzpgqutojoyzsnyfwsyeqqzrlhzbc"
	fmt.Println(longestPalindrome(s))

}
