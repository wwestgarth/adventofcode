package solutions

import (
	utils "adventofcode/utils"
)

const (
	gTreeRepresentation = '#'
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
			mapData.treeData[yIndex][xIndex] = (point == gTreeRepresentation)
		}
	}

	return
}

func treesForSlope(data MapData, slope utils.Vector) (nTrees int) {

	position := utils.NewZeroVector()
	for {

		if data.isTreeAtPosition(position.X, position.Y) {
			nTrees += 1
		}

		position = utils.VecAdd(position, slope)
		if position.Y > data.length {
			break
		}
	}

	return nTrees

}

func SolveDay3() (part1, part2 string) {

	slopes := [5]utils.Vector{
		utils.NewVector(3, 1, 0),
		utils.NewVector(1, 1, 0),
		utils.NewVector(5, 1, 0),
		utils.NewVector(7, 1, 0),
		utils.NewVector(1, 2, 0),
	}

	data := fileAsMapData()
	treeCount := treesForSlope(data, slopes[0])

	treeCountMulti := 1

	for _, slope := range slopes {
		treeCountMulti *= treesForSlope(data, slope)
	}

	return utils.ResultStrings(treeCount, treeCountMulti)
}
