package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func readProgram(path string) (string, error) {
	buf, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

func part1(program string) (int, error) {
	r := regexp.MustCompile("mul\\((\\d+),(\\d+)\\)")
	matches := r.FindAllStringSubmatch(program, -1)

	var sum int
	for i := 0; i < len(matches); i += 1 {
		a, err := strconv.Atoi(matches[i][1])
		if err != nil {
			return 0, err
		}

		b, err := strconv.Atoi(matches[i][2])
		if err != nil {
			return 0, err
		}

		sum += a * b
	}
	return sum, nil
}

func part2(program string) (int, error) {
	var sum int

	mulReg := regexp.MustCompile("mul\\((\\d{1,3}),(\\d{1,3})\\)")
	doReg := regexp.MustCompile("do\\(\\)")
	dontReg := regexp.MustCompile("don't\\(\\)")

	var current int
	for current < len(program) {
		var currentDont int
		dontResult := dontReg.FindStringIndex(program[current:])
		if dontResult != nil {
			currentDont = current + dontResult[0]
		} else {

			currentDont = len(program)
		}

		matches := mulReg.FindAllStringSubmatch(program[current:currentDont], -1)
		for _, match := range matches {
			a, err := strconv.Atoi(match[1])
			if err != nil {
				return 0, err
			}

			b, err := strconv.Atoi(match[2])
			if err != nil {
				return 0, err
			}

			sum += a * b
		}

		doResult := doReg.FindStringIndex(program[currentDont:])
		if doResult != nil {
			current = currentDont + doResult[1]
		} else {
			current = len(program)
		}

	}
	return sum, nil
}

func main() {
	program, err := readProgram("./input.txt")
	if err != nil {
		panic(err)
	}

	result1, err := part1(program)
	if err != nil {
		panic(err)
	}

	result2, err := part2(program)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Result (part 1): %d\nResult (part 2): %d\n", result1, result2)
}
