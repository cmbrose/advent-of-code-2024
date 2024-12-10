package main

import (
	"fmt"
	"strings"

	"main/util"
)

func main() {
	reports := util.Map(util.ReadInputLines(), func(line string) []int {
		return util.Map(strings.Split(line, " "), util.AssertInt)
	})

	safe := util.Filter(reports, func(report []int) bool {
		sign := util.Sign(report[0] - report[1])

		for i := 0; i < len(report)-1; i += 1 {
			diff := report[i] - report[i+1]
			if sign != util.Sign(diff) {
				return false
			}
			abs := util.Abs(diff)
			if abs < 1 || abs > 3 {
				return false
			}
		}
		return true
	})

	fmt.Printf("%d\n", len(safe))
}
