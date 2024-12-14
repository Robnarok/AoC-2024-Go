package one

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

func First(input []string) {
	var lefts []int
	var rights []int
	for _, line := range input {
		left, right, err := readValues(line)
		if err != nil {
			log.Fatal(err)
		}
		lefts = append(lefts, left)
		rights = append(rights, right)
	}
	lefts, rights = sortNumbers(lefts, rights)
	var diffferences []int
	for i := range input {
		diffferences = append(diffferences, getDifference(lefts[i], rights[i]))
	}
	sum := 0
	for _, v := range diffferences {
		sum += v
	}
	fmt.Printf("Sum off differences: %v", sum)
}

func Second(input []string) {
	var lefts []int
	var rights []int
	for _, line := range input {
		left, right, err := readValues(line)
		if err != nil {
			log.Fatal(err)
		}
		lefts = append(lefts, left)
		rights = append(rights, right)
	}
	var simmilarities []int
	for i := range input {
		simmilarities = append(simmilarities, getSimilarity(lefts[i], rights))
	}
	sum := 0
	for i, v := range simmilarities {
		sum += lefts[i] * v
	}
	fmt.Printf("Sum off all simmilarities: %v", sum)
}

func getSimilarity(left int, rights []int) int {
	simmilarity := 0
	for _, v := range rights {
		if v == left {
			simmilarity++
		}
	}
	return simmilarity
}

func sortNumbers(lefts, rights []int) ([]int, []int) {
	sort.Ints(lefts)
	sort.Ints(rights)
	return lefts, rights
}

func readValues(line string) (int, int, error) {
	s := strings.Split(line, "   ")
	if len(s) != 2 {
		return 0, 0, fmt.Errorf("readValues(%v): got %d Values, not 2", line, len(s))
	}

	left, err := strconv.Atoi(s[0])
	if err != nil {
		return 0, 0, err
	}
	right, err := strconv.Atoi(s[1])
	if err != nil {
		return 0, 0, err
	}
	return left, right, nil
}

func getDifference(left, right int) int {
	difference := left - right

	if difference < 0 {
		difference *= -1
	}
	return difference
}
