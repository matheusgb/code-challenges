package main

import (
	"fmt"
	"sort"
	"strings"
)

func isAnagram(s string, t string) bool {
	ss := strings.Split(s, "")
	tt := strings.Split(t, "")
	sort.Strings(ss)
	sort.Strings(tt)

	return strings.Join(ss, "") == strings.Join(tt, "")
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
}
