package main

import (
	"os"
	"strconv"

	aoc2024go "github.com/Robnarok/AoC-2024-Go"
)

func main() {
	i, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}

	aoc2024go.Run(i)
}
