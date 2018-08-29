package main

import (
	"fmt"
)

func main() {

	strs := []string{
		"abcde",
		"sdaklfj",
		"asdjf",
		"na",
		"basdn",
		"sdaklfj",
		"asdjf",
		"na",
		"asdjf",
		"na",
		"basdn",
		"sdaklfj",
		"asdjf",
	}

	queries := []string{
		"abcde",
		"sdaklfj",
		"asdjf",
		"na",
		"basdn",
	}

	table := make(map[string]int)
	for _, each := range strs {
		table[each]++
	}
	var ret []int
	for _, each := range queries {
		count, ok := table[each]
		if ok {
			ret = append(ret, count)
		} else {
			ret = append(ret, 0)
		}
	}

	fmt.Println(ret)
}
