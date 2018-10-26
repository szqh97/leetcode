package main

import "fmt"

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	pali_str := ""
	for i := 0; i < len(s); i++ {
		for j := 0; j <= i && i+j < len(s); j++ {
			if i-j >= 0 {
				tmps := string(s[i-j : i+j+1])
				if tmps == reverse(tmps) && len(tmps) > len(pali_str) {
					pali_str = tmps
				}
			}
			if i-j-1 >= 0 {
				tmps2 := string(s[i-j-1 : i+j+1])

				if tmps2 == reverse(tmps2) && len(tmps2) > len(pali_str) {
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
