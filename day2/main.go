package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

type AnalisysResult struct {
	beforeIncreasingTolerance int
	afterIncreasingTolerance  int
}

func readReports(path string) ([][]int, error) {
	buf, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(buf), "\n")

	var reports [][]int
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		levels := strings.Split(line, " ")

		var report []int
		for _, level := range levels {
			levelNum, err := strconv.Atoi(level)
			if err != nil {
				return nil, err
			}
			report = append(report, levelNum)
		}
		reports = append(reports, report)
	}
	return reports, nil
}

func checkSafety(a, b int, isIncreasing bool) bool {
	difference := a - b
	absDifference := abs(difference)

	if absDifference <= 0 || absDifference > 3 {
		return false
	}
	if isIncreasing {
		if difference >= 0 {
			return false
		}
	} else {
		if difference <= 0 {
			return false
		}
	}
	return true
}

func checkReport(report []int) bool {
	var isIncreasing bool
	if report[0] < report[1] {
		isIncreasing = true
	} else if report[0] > report[1] {
		isIncreasing = false
	}
	if !checkSafety(report[0], report[1], isIncreasing) {
		return false
	}

	for i := 1; i < len(report)-1; i += 1 {
		if !checkSafety(report[i], report[i+1], isIncreasing) {
			return false
		}
	}
	return true
}

func checkTolerance(rejectedReport chan []int) int {
	var counter int
	for {
		report, ok := <-rejectedReport
		if !ok {
			break
		}

		for removedNum := 0; removedNum < len(report); removedNum += 1 {
			newReport := append([]int(nil), report[:removedNum]...)
			newReport = append(newReport, report[removedNum+1:]...)

			if checkReport(newReport) {
				counter += 1
				break
			}
		}
	}
	return counter
}

func countSafeReports(reports [][]int, counter *int, rejectedReport chan []int) {
	for _, report := range reports {
		if checkReport(report) {
			*counter += 1
		} else {
			rejectedReport <- report
		}
	}
	close(rejectedReport)
}

func part1AndPart2(reports [][]int) AnalisysResult {
	rejectedReport := make(chan []int)

	var counter int
	go countSafeReports(reports, &counter, rejectedReport)

	result := checkTolerance(rejectedReport)

	return AnalisysResult{counter, counter + result}
}

func main() {
	reports, err := readReports("./input.txt")
	if err != nil {
		panic(err)
	}
	result := part1AndPart2(reports)
	fmt.Printf("Result (part 1): %d\nResult (part 2): %d\n", result.beforeIncreasingTolerance, result.afterIncreasingTolerance)
}
