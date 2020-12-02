package solutions

import (
	utils "adventofcode/utils"
	"strconv"
	"strings"
)

type PolicyData struct {
	bounds    [2]int
	character string
	password  string
}

func (data *PolicyData) part1Valid() bool {

	count := strings.Count(data.password, data.character)
	return count >= data.bound(0) && count <= data.bound(1)
}

func (data *PolicyData) part2Valid() bool {
	l, u := data.AtIndices()

	return (l == data.character) != (u == data.character)

}

func (data *PolicyData) AtIndices() (lower, upper string) {
	indexL := data.bound(0) - 1
	indexU := data.bound(1) - 1
	lower, upper = "", ""

	if indexL < len(data.password) {
		lower = data.password[indexL : indexL+1]
	}

	if indexU < len(data.password) {
		upper = data.password[indexU : indexU+1]
	}
	return
}

func (data *PolicyData) bound(i int) int {
	return data.bounds[i]
}

func NewPolicyData(input string) PolicyData {

	var data PolicyData

	splitInput := strings.Fields(input)
	splitBound := strings.Split(splitInput[0], "-")

	data.bounds[0], _ = strconv.Atoi(splitBound[0])
	data.bounds[1], _ = strconv.Atoi(splitBound[1])
	data.character = splitInput[1][:1]
	data.password = splitInput[2]

	return data

}

func transformInput2() ([]string, error) {

	return utils.ReadFileAsStrings("input/day2.txt")
}

func transformOutput2(value1, value2 int, err error) (part1Result, part2Result string) {

	if err != nil {
		return "", ""
	}

	return strconv.Itoa(value1), strconv.Itoa(value2)
}

func SolveDay2() (string, string) {

	inputs, err := transformInput2()
	part1Valids := 0
	part2Valids := 0

	for _, input := range inputs {

		data := NewPolicyData(input)

		if !data.part1Valid() {
			part1Valids += 1
		}
		if !data.part2Valid() {
			part2Valids += 1
		}
	}

	return transformOutput2(part1Valids, len(inputs)-part2Valids, err)
}
