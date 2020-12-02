package solutions

import (
	utils "adventofcode/utils"
)

var (
	gSum = 2020
)

func SolveDay1() (part1, part2 string) {

	tracker := make(map[int]int)
	part1Result, part2Result := 0, 0

	inputs := utils.FileAsInts("input/day1.txt")

	for _, input := range inputs {

		if compliment, in := tracker[input]; in {
			part1Result = compliment * input
			break
		}
		tracker[gSum-input] = input
	}

	return utils.ResultStrings(part1Result, part2Result)
}
