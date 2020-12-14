package solutions

import (
	utils "adventofcode/utils"
	"fmt"
)

var (
	North = utils.NewVector(0, 1, 0)
	South = utils.NewVector(0, -1, 0)
	East  = utils.NewVector(1, 0, 0)
	West  = utils.NewVector(-1, 0, 0)

	enumeratedDirections = map[int]utils.Vector{0: North, 1: East, 2: South, 3: West}
)

type wayStation struct {
	position utils.Vector
}

func NewWayStation() (w wayStation) {
	w.position = utils.NewVector(10, 1, 0)
	return w
}

func (w *wayStation) Move(d utils.Vector, v int) {

	scaled := utils.VecMultiply(d, v)
	w.position = utils.VecAdd(w.position, scaled)
}

func (w *wayStation) Turn(v int) {

	for i := 0; i < v/90; i++ {
		w.position = utils.NewVector(w.position.Y, -w.position.X, 0)
	}
}

type ship struct {

	// Angles look to be all multiples of 90
	facing   int
	position utils.Vector
}

func NewShip() (s ship) {

	s.facing = 1
	s.position = utils.NewVector(0, 0, 0)
	return
}

func (s *ship) Move(d utils.Vector, v int) {

	scaled := utils.VecMultiply(d, v)
	s.position = utils.VecAdd(s.position, scaled)
}

func (s *ship) Forward(v int) {

	scaled := utils.VecMultiply(s.Facing(), v)
	s.position = utils.VecAdd(s.position, scaled)
}

func (s *ship) Turn(v int) {

	s.facing = (s.facing + (v / 90)) % 4
}

func (s *ship) Position() utils.Vector {
	return s.position
}

func (s *ship) Facing() utils.Vector {
	return enumeratedDirections[s.facing]
}

type vessel interface {
	Turn(int)
	Move(utils.Vector, int)
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func instruct(v vessel, instruction rune, value int) {
	if instruction == 'R' {
		v.Turn(value)
	} else if instruction == 'L' {
		v.Turn(360 - value)
	} else if instruction == 'N' {
		v.Move(North, value)
	} else if instruction == 'S' {
		v.Move(South, value)
	} else if instruction == 'E' {
		v.Move(East, value)
	} else if instruction == 'W' {
		v.Move(West, value)
	}
}

func solvePart1() int {
	ship := NewShip()

	for _, line := range utils.FileAsLines("input/day12.txt") {

		instruction := line[0]
		value := utils.AtoiPanic(line[1:])

		if instruction == 'F' {
			ship.Forward(value)
		} else {
			instruct(&ship, rune(instruction), value)
		}

	}

	fmt.Println(ship.position.X, ship.position.Y)
	return Abs(ship.position.X) + Abs(ship.position.Y)
}

func solvePart2() int {
	ship := NewShip()
	w := NewWayStation()

	for _, line := range utils.FileAsLines("input/day12.txt") {

		instruction := line[0]
		value := utils.AtoiPanic(line[1:])

		if instruction == 'F' {
			scaled := utils.VecMultiply(w.position, value)
			ship.position = utils.VecAdd(ship.position, scaled)
		} else {
			instruct(&w, rune(instruction), value)
		}

	}

	fmt.Println(ship.position.X, ship.position.Y)
	return Abs(ship.position.X) + Abs(ship.position.Y)
}

func SolveDay12() (string, string) {

	part1 := solvePart1()
	part2 := solvePart2()

	return utils.ResultStrings(part1, part2)
}
