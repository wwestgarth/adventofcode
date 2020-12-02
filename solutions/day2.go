package solutions

import (
	utils "adventofcode/utils"
	"strings"
)

type PolicyData struct {
	bounds    [2]int
	character string
	password  string
}

func (data *PolicyData) validPolicy1() bool {

	count := strings.Count(data.password, data.character)
	return count >= data.bound(0) && count <= data.bound(1)
}

func (data *PolicyData) validPolicy2() bool {
	l, u := data.AtIndices()

	return (l == data.character) != (u == data.character)

}

func (data *PolicyData) AtIndices() (lower, upper string) {
	indexL := data.bound(0) - 1
	indexU := data.bound(1) - 1

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

func NewPolicyData(input string) (data PolicyData) {

	splitInput := strings.Fields(input)
	splitBound := strings.Split(splitInput[0], "-")

	data.bounds[0] = utils.AtoiPanic(splitBound[0])
	data.bounds[1] = utils.AtoiPanic(splitBound[1])

	data.character = splitInput[1][:1]
	data.password = splitInput[2]

	return data

}

func fileAsPolicyData() (policyData []PolicyData) {

	inputs := utils.FileAsLines("input/day2.txt")

	for _, input := range inputs {
		policyData = append(policyData, NewPolicyData(input))
	}

	return
}

func SolveDay2() (part1, part2 string) {

	policyData := fileAsPolicyData()

	validPolicy1 := 0
	invalidPolicy2 := len(policyData)

	for _, data := range policyData {

		if !data.validPolicy1() {
			validPolicy1 += 1
		}
		if !data.validPolicy2() {
			invalidPolicy2 -= 1
		}
	}

	return utils.ResultStrings(validPolicy1, invalidPolicy2)
}
