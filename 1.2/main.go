package main

import (
	"fmt"
	"main/util"
	"strings"
)

func main() {
	left := make(map[int]int)
	right := make(map[int]int)

	for _, line := range util.ReadInputLines() {
		parts := strings.Split(line, " ")
		l := util.AssertInt(parts[0])
		left[l] += 1

		r := util.AssertInt(parts[len(parts)-1])
		right[r] += 1
	}

	sum := 0
	for k, lcnt := range left {
		rcnt := right[k]
		sum += k * lcnt * rcnt
	}

	fmt.Printf("%d\n", sum)
}
