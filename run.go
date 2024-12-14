package aoc2024go

import (
	"log"

	"github.com/Robnarok/AoC-2024-Go/days/one"
	"github.com/Robnarok/AoC-2024-Go/helper"
)

func Run(day int) {
	if day == 1 {
		input := helper.Parse("input/1.txt")
		log.Println("Day 1: Part One")
		one.First(input)
		log.Println("Day 1: Part Two")
		one.Second(input)
	}
}
