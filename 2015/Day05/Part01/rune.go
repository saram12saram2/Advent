package main

import (
	"fmt"
	"strings"
)

func isVowel(ch rune) bool {
	return strings.ContainsRune("aeiou", ch)
}

func isNice(s string) bool {
	vowelCount, doubleLetter, containsBadStr := 0, false, false

	// Check for disallowed strings
	for _, badStr := range []string{"ab", "cd", "pq", "xy"} {
		if strings.Contains(s, badStr) {
			containsBadStr = true
			break
		}
	}

	// Check for vowels and double letters
	for i, ch := range s {
		if isVowel(ch) {
			vowelCount++
		}
		if i > 0 && ch == rune(s[i-1]) {
			doubleLetter = true
		}
	}

	return vowelCount >= 3 && doubleLetter && !containsBadStr
}

func main() {
	strings := []string{
		"uxcplgxnkwbdwhrp", "suerykeptdsutidb", "dmrtgdkaimrrwmej", "ztxhjwllrckhakut",
		"gdnzurjbbwmgayrg", "gjdzbtrcxwprtery", "fbuqqaatackrvemm", "pcjhsshoveaodyko",
		"lrpprussbesniilv", "mmsebhtqqjiqrusd", "vumllmrrdjgktmnb", "ptsqjcfbmgwdywgi",
		"mmppavyjgcfebgpl", "zexyxksqrqyonhui", "npulalteaztqqnrl", "mscqpccetkktaknl",
		"ydssjjlfejdxrztr", "jdygsbqimbxljuue", "ortsthjkmlonvgci", "jfjhsbxeorhgmstc",
		"vdrqdpojfuubjbbg", "xxxddetvrlpzsfpq", "zpjxvrmaorjpwegy", "laxrlkntrukjcswz",
		"pbqoungonelthcke", "niexeyzvrtrlgfzw", "zuetendekblknqng", "lyazavyoweyuvfye",
		"tegbldtkagfwlerf", "xckozymymezzarpy",
	}

	niceCount := 0
	for _, s := range strings {
		if isNice(s) {
			niceCount++
		}
	}

	fmt.Printf("Number of nice strings: %d\n", niceCount)
}
