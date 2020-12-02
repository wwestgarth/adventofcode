package solutions

import (
	utils "adventofcode/utils"
	"strconv"
	"strings"
)

type Data struct {
	lowerBound int
	upperBound int
	character  string
	password   string
}

func (data *Data) invalid() bool {

	count := strings.Count(data.password, data.character)
	if count > data.upperBound || count < data.lowerBound {
		return false
	}

	return true
}

func resolveInput(input string) Data {

	splitInput := strings.Fields(input)
	splitBound := strings.Split(splitInput[0], "-")

	var data Data
	data.lowerBound, _ = strconv.Atoi(splitBound[0])
	data.upperBound, _ = strconv.Atoi(splitBound[1])
	data.character = splitInput[1][:1]
	data.password = splitInput[2]

	return data

}

func SolveDay2() string {

	inputs, _ := utils.ReadFileAsStrings("input/day2.txt")
	bads := 0

	for _, input := range inputs {

		data := resolveInput(input)

		if data.invalid() {
			bads += 1
		}
	}

	return strconv.Itoa(bads)
}
