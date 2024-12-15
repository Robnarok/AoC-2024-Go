package aoc2024go

import (
	"fmt"
	"log"

	"github.com/Robnarok/AoC-2024-Go/days/one"
	"github.com/Robnarok/AoC-2024-Go/days/three"
	"github.com/Robnarok/AoC-2024-Go/days/two"
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
	if day == 2 {
		input := helper.Parse("input/2.txt")
		log.Println("Day 2: Part One")
		fmt.Println(two.First(input))
		log.Println("Day 2: Part Two")
		fmt.Println(two.Second(input))
	}
	if day == 3 {
		input := helper.Parse("input/3.txt")
		log.Println("Day 3: Part One")
		fmt.Println(three.First(input))
		log.Println("Day 3: Part Two")
		fmt.Println(three.Second(input))
	}
}
