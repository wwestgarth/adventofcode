package solutions

import (
	utils "adventofcode/utils"
	"fmt"
	"sort"
)

const (
	maxRows    = 128
	maxColumns = 8
)

func calculateSeatId(seatStr string) (int, int, int) {

	rowPattern := seatStr[:7]
	columnPattern := seatStr[7:]
	//columnId := 0;
	rowUpper := maxRows
	rowLower := 1
	columnUpper := maxColumns
	columnLower := 1

	for _, row := range rowPattern {

		mid := (rowUpper + rowLower) / 2
		if row == 'F' {
			rowUpper = mid
		} else {
			rowLower = mid
		}
	}

	for _, column := range columnPattern {

		mid := (columnUpper + columnLower) / 2
		if column == 'L' {
			columnUpper = mid
		} else {
			columnLower = mid
		}
	}

	return (rowLower * 8) + columnLower, rowLower, columnLower
}

func SolveDay5() (string, string) {

	part1 := 0
	part2 := 0

	seatInfo := utils.FileAsLines("input/day5.txt")
	allKnownSeats := make([]int, 0)
	seatRow := make(map[int]int)
	seatColumn := make(map[int]int)
	for _, info := range seatInfo {
		id, r, c := calculateSeatId(info)
		allKnownSeats = append(allKnownSeats, id)
		seatRow[id] = r
		seatColumn[id] = c
		if id > part1 {
			part1 = id
		}

	}

	sort.Ints(allKnownSeats)
	for i, v := range allKnownSeats {

		if i == 0 {
			continue
		}

		prev := allKnownSeats[i-1]
		if v-prev == 2 && seatColumn[prev] != 7 && seatColumn[v] != 1 {
			fmt.Println(seatRow[prev], seatColumn[prev])
			fmt.Println(seatRow[v], seatColumn[v])
			part2 = v + 1
		}

	}

	return utils.ResultStrings(part1, part2)
}
