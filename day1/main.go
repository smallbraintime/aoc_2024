package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
)

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

type LocationList struct {
	left, right []int
}

func readLocations(path string) (LocationList, error) {
	buf, err := os.ReadFile(path)
	if err != nil {
		return LocationList{nil, nil}, err
	}
	lines := strings.Split(string(buf), "\n")

	var left, right []int
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		words := strings.Split(line, "   ")

		ln, err := strconv.Atoi(words[0])
		if err != nil {
			return LocationList{nil, nil}, err
		}

		rn, err := strconv.Atoi(words[1])
		if err != nil {
			return LocationList{nil, nil}, err
		}

		left = append(left, ln)
		right = append(right, rn)
	}
	return LocationList{left, right}, nil
}

func sortLocations(list LocationList) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		slices.Sort(list.left)
	}()

	go func() {
		defer wg.Done()
		slices.Sort(list.right)
	}()

	wg.Wait()
}

func sumLocationDistances(list LocationList) int {
	var sum int
	for i := 0; i < len(list.left); i++ {
		sum += int(abs(list.left[i] - list.right[i]))
	}
	return sum
}

func countBothOccurences(list LocationList) map[int]int {
	leftOccurences := make(map[int]int)
	rightOccurences := make(map[int]int)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for _, v := range list.left {
			leftOccurences[v] += 1
		}
	}()

	go func() {
		defer wg.Done()
		for _, v := range list.right {
			rightOccurences[v] += 1
		}
	}()

	wg.Wait()

	var occurences = make(map[int]int)
	for k, lv := range leftOccurences {
		rv, ok := rightOccurences[k]
		if ok {
			occurences[k] = max(lv, rv)
		}
	}
	return occurences
}

func calcSimilarityScore(occurences map[int]int) int {
	var score int
	for k, v := range occurences {
		score += k * v
	}
	return score
}

func part1(data LocationList) int {
	sortLocations(data)
	return sumLocationDistances(data)
}

func part2(data LocationList) int {
	return calcSimilarityScore(countBothOccurences(data))
}

func main() {
	data, err := readLocations("./input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Result (part 1): %d\nResult (part2): %d\n", part1(data), part2(data))
}
