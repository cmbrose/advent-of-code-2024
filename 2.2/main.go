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
		return isReportValid(report) || bruteForce(report)
	})

	fmt.Printf("%d\n", len(safe))
}

func bruteForce(report []int) bool {
	for i := range report {
		var reportWithout []int
		reportWithout = append(reportWithout, report[:i]...)
		reportWithout = append(reportWithout, report[i+1:]...)

		if isReportValid(reportWithout) {
			return true
		}
	}
	return false
}

func isReportValid(report []int) bool {
	sign := util.Sign(report[0] - report[1])

	for i := 0; i < len(report)-1; i += 1 {
		if !isPairValid(report[i], report[i+1], sign) {
			return false
		}
	}
	return true
}

func isPairValid(a, b, s int) bool {
	diff := a - b
	if s != util.Sign(diff) {
		return false
	}

	abs := util.Abs(diff)
	if abs < 1 || abs > 3 {
		return false
	}

	return true
}
