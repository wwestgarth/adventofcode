package solutions

import (
	utils "adventofcode/utils"
	"fmt"
	"log"
	"strings"
	"time"
)

func fileAsGameInput() []int {

	var inputs []int
	input := utils.FileAsLines("input/day15.txt")

	for _, i := range strings.Split(input[0], ",") {
		inputs = append(inputs, utils.AtoiPanic(i))
	}
	return inputs

}

func addSpokenNumber(spoken map[int][]int, turn, num int) bool {

	wasNew := false
	_, ok := spoken[num]

	if !ok {
		wasNew = true
		spoken[num] = make([]int, 0)
	}

	if len(spoken[num]) == 2 {
		spoken[num][0] = spoken[num][1]
		spoken[num][1] = turn
	} else {
		spoken[num] = append(spoken[num], turn)
	}
	return wasNew
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func SolveDay15() (string, string) {
	defer timeTrack(time.Now(), "Solve")
	part1 := 0
	part2 := 0

	spoken := make(map[int][]int)
	inputs := fileAsGameInput()

	// Get started

	for i, num := range inputs {
		wasNew := addSpokenNumber(spoken, i+1, num)
		fmt.Println(i+1, ":", num, wasNew, spoken[num])
	}

	// Add first again
	wasNew := addSpokenNumber(spoken, len(inputs)+1, 0)

	turn := len(spoken) + 1
	currentNumber := 0
	firstNumber := 0
	for turn < 30000000 {

		if wasNew {
			currentNumber = firstNumber
		} else {
			// Find next number
			calls := spoken[currentNumber]

			currentNumber = turn - calls[len(calls)-2]
		}

		turn += 1
		wasNew = addSpokenNumber(spoken, turn, currentNumber)
		//fmt.Println(turn, ":", currentNumber, wasNew) // spoken[currentNumber])

	}
	part1 = currentNumber
	return utils.ResultStrings(part1, part2)
}
