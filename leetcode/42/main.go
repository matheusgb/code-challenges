package main

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	a := [][]string{}

	for _, str := range strs {
		ss := strings.Split(str, "")
		sort.Strings(ss)
		jss := strings.Join(ss, "")
		if m[jss] == nil {
			m[jss] = []string{str}
		} else {
			m[jss] = append(m[jss], str)
		}
	}

	for _, mm := range m {
		a = append(a, mm)
	}

	return a
}

func main() {
	fmt.Println(groupAnagrams([]string{"", "tea", "tan", "ate", "nat", "bat"}))
}
