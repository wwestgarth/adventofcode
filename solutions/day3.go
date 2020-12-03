package solutions

import (
	utils "adventofcode/utils"
)

type MapData struct {
	width    int
	length   int
	treeData map[int]map[int]bool
}

func (data *MapData) isTreeAtPosition(x, y int) bool {

	wrappedX := x % data.width
	return data.treeData[y][wrappedX]
}

func fileAsMapData() (mapData MapData) {

	inputs := utils.FileAsLines("input/day3.txt")

	mapData.treeData = make(map[int]map[int]bool)
	mapData.length = len(inputs)
	mapData.width = len(inputs[0])

	for yIndex, line := range inputs {

		mapData.treeData[yIndex] = make(map[int]bool)

		for xIndex, point := range line {
			mapData.treeData[yIndex][xIndex] = (point == '#')
		}
	}

	return
}

func step1(x, y int) (int, int) {
	return x + 3, y + 1
}

func step2(x, y int) (int, int) {
	return x + 1, y + 1
}

func step3(x, y int) (int, int) {
	return x + 5, y + 1
}

func step4(x, y int) (int, int) {
	return x + 7, y + 1
}

func step5(x, y int) (int, int) {
	return x + 1, y + 2
}

func SolveDay3() (part1, part2 string) {

	x := 0
	y := 0
	treeCount1 := 0
	treeCount2 := 0
	treeCount3 := 0
	treeCount4 := 0
	treeCount5 := 0

	data := fileAsMapData()
	for {

		if data.isTreeAtPosition(x, y) {
			treeCount1 += 1
		}

		x, y = step1(x, y)
		if y > data.length {
			break
		}
	}
	x = 0
	y = 0
	for {

		if data.isTreeAtPosition(x, y) {
			treeCount2 += 1
		}

		x, y = step2(x, y)
		if y > data.length {
			break
		}
	}
	x = 0
	y = 0
	for {

		if data.isTreeAtPosition(x, y) {
			treeCount3 += 1
		}

		x, y = step3(x, y)
		if y > data.length {
			break
		}
	}
	x = 0
	y = 0
	for {

		if data.isTreeAtPosition(x, y) {
			treeCount4 += 1
		}

		x, y = step4(x, y)
		if y > data.length {
			break
		}
	}
	x = 0
	y = 0
	for {

		if data.isTreeAtPosition(x, y) {
			treeCount5 += 1
		}

		x, y = step5(x, y)
		if y > data.length {
			break
		}
	}

	return utils.ResultStrings(treeCount1, treeCount1*treeCount2*treeCount3*treeCount4*treeCount5)
}
