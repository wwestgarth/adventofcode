package solutions

import (
	utils "adventofcode/utils"
	"strconv"
)

var (
	gSum = 2020
)

func transformInput1() ([]int, error) {
	return utils.ReadFileAsInts("input/day1.txt")
}

func transformOutput1(value1, value2 int, err error) (part1Result, part2Result string) {

	if err != nil {
		return "", ""
	}

	return strconv.Itoa(value1), strconv.Itoa(value2)
}

func SolveDay1() (string, string) {

	tracker := make(map[int]int)
	part1Result, part2Result := 0, 0

	inputs, err := transformInput1()

	for _, input := range inputs {

		if compliment, in := tracker[input]; in {
			part1Result = compliment * input
			break
		}
		tracker[gSum-input] = input
	}

	return transformOutput1(part1Result, part2Result, err)
}
