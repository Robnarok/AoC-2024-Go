package three

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func First(input []string) int {
	ret := 0
	for _, line := range input {
		s := filterLine(line)
		for _, elemtent := range s {
			i, i2, err := parseCommand(elemtent)
			if err != nil {
				log.Fatal(err)
			}
			tmp := i * i2
			ret += tmp
		}
	}
	return ret
}

// Lines need to be merged into one Line, because the Disable/Enable is persistent between lines :D
func Second(input []string) int {
	ret := 0
	mergedLines := ""
	for _, line := range input {
		mergedLines += line
	}
	mergedLines = prepareLine(mergedLines)
	s := filterLine(mergedLines)
	for _, elemtent := range s {
		i, i2, err := parseCommand(elemtent)
		if err != nil {
			log.Fatal(err)
		}
		tmp := i * i2
		ret += tmp
	}
	return ret
}

// .*? to prevent greedy matches
func prepareLine(input string) string {
	muster := `don't\(\).*?do\(\)`
	m := regexp.MustCompile(muster)
	output := m.FindAllString(input, -1)
	for _, match := range output {
		input = strings.ReplaceAll(input, match, "AAAA")
	}

	// Replace the last Don't
	muster = `don't\(\).*`
	m = regexp.MustCompile(muster)
	output = m.FindAllString(input, -1)
	for _, match := range output {
		fmt.Println("It happend!")
		input = strings.ReplaceAll(input, match, "")
	}

	return input
}

func filterLine(input string) []string {
	muster := `mul\([0-9]+,[0-9]+\)`
	m := regexp.MustCompile(muster)
	output := m.FindAllString(input, -1)
	return output
}

func parseCommand(input string) (int, int, error) {
	s := strings.Trim(input, "mul()")
	s2 := strings.Split(s, ",")

	if len(s2) != 2 {
		return 0, 0, fmt.Errorf("parseCommand: %v", "Ungleich 2 Elemente im array")
	}

	i, err := strconv.Atoi(s2[0])
	if err != nil {
		return 0, 0, err
	}
	j, err := strconv.Atoi(s2[1])
	if err != nil {
		return 0, 0, err
	}

	return i, j, nil
}
