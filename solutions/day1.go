package solutions

import (
	utils "adventofcode/utils"
	"strconv"
)

func SolveDay1() string {

	sum := 2020
	inputs, _ := utils.ReadFileAsInts("input/day1.txt")

	tracker := make(map[int]int)

	for _, input := range inputs {

		value, ok := tracker[input]
		if ok {
			return strconv.Itoa(value * input)
		}

		tracker[sum-input] = input
	}

	return ""
}
