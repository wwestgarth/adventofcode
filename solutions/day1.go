package day1

import (
	utils "adventofcode/utils"
	"strconv"
)

func Solve() string {

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
