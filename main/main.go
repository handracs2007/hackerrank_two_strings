// https://www.hackerrank.com/challenges/two-strings/problem

package main

import "fmt"

func twoStrings(s1 string, s2 string) string {
	var charMap = make(map[uint8]int)

	// Ler's map the characters first
	for i := 0; i < len(s1); i++ {
		charMap[s1[i]] = charMap[s1[i]] + 1
	}

	// Now let's check the substring
	for i := 0; i < len(s2); i++ {
		if charMap[s2[i]] > 0 {
			return "YES"
		}
	}

	return "NO"
}

func main() {
	fmt.Println(twoStrings("hello", "world"))
	fmt.Println(twoStrings("hi", "world"))
}
