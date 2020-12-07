package solutions

import (
	utils "adventofcode/utils"
	"strings"

	"github.com/golang-collections/collections/set"
)

func setFromAnswers(answers string) *set.Set {

	s := set.New()
	for _, t := range answers {
		s.Insert(t)
	}

	return s
}

func SolveDay6() (string, string) {

	part1 := 0
	part2 := 0

	inputs := utils.FileAsLines("input/day6.txt")

	data := ""
	for _, input := range inputs {

		if input == "" {

			uniqueAnswers := set.New()
			for _, answer := range data {
				uniqueAnswers.Insert(answer)
			}

			part1 += uniqueAnswers.Len()
			data = ""
		}
		data += input
	}

	data = ""
	for _, input := range inputs {

		if input == "" {

			groupAnswers := strings.Fields(data)

			all := setFromAnswers(groupAnswers[0])

			for i, answer := range groupAnswers {

				if i == 0 {
					continue
				}

				tSet := setFromAnswers(answer)
				all = all.Intersection(tSet)

			}

			part2 += all.Len()
			data = ""
		}
		data += " "
		data += input
	}

	return utils.ResultStrings(part1, part2)
}
