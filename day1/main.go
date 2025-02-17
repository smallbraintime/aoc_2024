package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
)

type LocationList struct {
	left, right []int
}

func readLocations(path string) (*LocationList, error) {
	buf, err := os.ReadFile(path)
	if err != nil {
		return nil, err
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
			return nil, err
		}

		rn, err := strconv.Atoi(words[1])
		if err != nil {
			return nil, err
		}

		left = append(left, ln)
		right = append(right, rn)
	}
	return &LocationList{left, right}, nil
}

func sortLocations(list *LocationList) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func(left []int) {
		defer wg.Done()
		slices.Sort(left)
	}(list.left)

	go func(right []int) {
		defer wg.Done()
		slices.Sort(right)
	}(list.right)

	wg.Wait()
}

func sumLocationDistances(list *LocationList) int {
	var sum int
	for i := 0; i < len(list.left); i++ {
		sum += int(math.Abs(float64(list.left[i] - list.right[i])))
	}
	return sum
}

func main() {
	data, err := readLocations("./input.txt")
	if err != nil {
		panic(err)
	}
	sortLocations(data)
	result := sumLocationDistances(data)
	fmt.Println("Result: ", result)
}
