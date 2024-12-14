package two

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func First(input []string) int {
	safetyCounter := 0
	for _, line := range input {
		i, err := parseLine(line)
		if err != nil {
			log.Fatal(err)
		}
		if checkSafety(i) {
			safetyCounter++
		}
	}
	return safetyCounter
}

func Second(input []string) int {
	safetyCounter := 0
	for _, line := range input {
		i, err := parseLine(line)
		if err != nil {
			log.Fatal(err)
		}
		if !checkSafety(i) {
			for j := range i {

				// Create shallow Copy of the Slice to modify it for the Damper
				tmp := make([]int, len(i))
				copy(tmp, i)
				tmp = remove(tmp, j)

				if checkSafety(tmp) {
					safetyCounter++
					break
				}
			}
		} else {
			safetyCounter++
		}
	}
	return safetyCounter
}

// Useless Remove function stolen from StackOverflow to remove an element from slice
func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func parseLine(line string) ([]int, error) {
	s := strings.Split(line, " ")
	var ret []int
	for i := range len(s) {
		tmp, err := strconv.Atoi(s[i])
		if err != nil {
			return nil, fmt.Errorf("parseLine: %v", err)
		}
		ret = append(ret, tmp)
	}
	return ret, nil
}

func checkSafety(data []int) bool {
	increase := 0
	decrease := 0
	for i := 0; i < len(data)-1; i++ {
		if data[i] > data[i+1] {
			if data[i]-data[i+1] > 3 {
				return false
			}
			decrease++
		}
		if data[i] < data[i+1] {
			if data[i+1]-data[i] > 3 {
				return false
			}
			increase++
		}
		if data[i] == data[i+1] {
			return false
		}
	}
	if increase == 0 || decrease == 0 {
		return true
	}
	return false
}
