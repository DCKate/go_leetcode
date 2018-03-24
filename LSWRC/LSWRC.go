//Package LSWRC porpose : find the Longest Substring Without Repeating Characters
package LSWRC

import (
	"fmt"
)

func isNotReapet(d map[rune]int, c rune) bool {
	if _, ok := d[c]; ok {
		return false
	}
	return true
}

func resetCheckArray(a []rune, d map[rune]int, now rune) ([]rune, map[rune]int) {
	ii, _ := d[now]
	ref, _ := d[a[0]]
	to := ii - ref + 1
	for j := 0; j < to; j++ {
		delete(d, a[j])
	}
	reset := a[to:]
	reset = append(reset, now)
	return reset, d
}

//LengthOfLongestSubstring :Given a string, find the length of the longest substring without repeating characters.
func LengthOfLongestSubstring(s string) string {
	tmpdict := make(map[rune]int)
	var longword []rune
	var tmpword []rune
	for i, c := range s {
		fmt.Printf("%v %c\n", i, c)
		if len(tmpword) == 0 || isNotReapet(tmpdict, c) {
			tmpword = append(tmpword, c)
		} else {
			if len(longword) < len(tmpword) {
				longword = tmpword[:]
			}
			fmt.Printf("1. %v\n", tmpword)
			tmpword, tmpdict = resetCheckArray(tmpword, tmpdict, c)
			fmt.Printf("2. %v\n", tmpword)
		}
		tmpdict[c] = i
	}

	if len(longword) < len(tmpword) {
		longword = tmpword[:]
	}
	return string(longword)
}

//LengthOfLongestSubstringv2 :Given a string, find the length of the longest substring without repeating characters.
func LengthOfLongestSubstringv2(s string) int {
	tmpdict := make(map[rune]int)
	tmprefer := 0
	longlen := 0
	for i, c := range s {
		if vv, ok := tmpdict[c]; ok {
			bef := i - tmprefer
			if longlen < bef {
				longlen = bef
			}
			if tmprefer <= vv {
				tmprefer = vv + 1
			}
			tmpdict[c] = i
			// fmt.Printf("%c : %v %v %v\n", c, longlen, tmprefer, tmpdict)
		} else {
			tmpdict[c] = i
		}
		if i == len(s)-1 && (tmpdict[c]-tmprefer) >= longlen {
			longlen = (tmpdict[c] - tmprefer) + 1
		}
	}
	return longlen
}
