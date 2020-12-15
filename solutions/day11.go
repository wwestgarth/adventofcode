package solutions

import (
	utils "adventofcode/utils"
	"fmt"
)

type state int

const (
	Unset = iota
	Floor
	Empty
	Occupied
)

type seatPlan struct {
	seats  [][]state
	width  int
	length int
}

func NewSeatPlan(width, length int) (s *seatPlan) {

	s = new(seatPlan)
	s.width = width
	s.length = length
	s.seats = make([][]state, length)
	for i := range s.seats {
		s.seats[i] = make([]state, width)
	}
	return s
}

func (s *seatPlan) state(x, y int) state {

	if x >= s.width || x < 0 {
		return Unset
	}

	if y >= s.length || y < 0 {
		return Unset
	}

	return s.seats[y][x]
}

func (s *seatPlan) setState(x, y int, newState state) {
	s.seats[y][x] = newState
}

func nextSeatInDir(s *seatPlan, x, y, v, w int) state {

	adjX := x
	adjY := y
	for {
		adjX += v
		adjY += w
		state := s.state(adjX, adjY)

		if state == Floor {
			continue
		}

		return state
	}
}

func nextState(s *seatPlan, x, y int) (newState state) {

	occupied := 0
	current := s.state(x, y)

	if current == Floor || current == Unset {
		return current
	}

	if nextSeatInDir(s, x, y, -1, -1) == Occupied {
		occupied += 1
	}
	if nextSeatInDir(s, x, y, 0, -1) == Occupied {
		occupied += 1
	}
	if nextSeatInDir(s, x, y, 1, -1) == Occupied {
		occupied += 1
	}
	if nextSeatInDir(s, x, y, -1, 0) == Occupied {
		occupied += 1
	}
	if nextSeatInDir(s, x, y, 1, 0) == Occupied {
		occupied += 1
	}
	if nextSeatInDir(s, x, y, -1, 1) == Occupied {
		occupied += 1
	}
	if nextSeatInDir(s, x, y, 0, 1) == Occupied {
		occupied += 1
	}
	if nextSeatInDir(s, x, y, 1, 1) == Occupied {
		occupied += 1
	}

	if occupied == 0 {
		return Occupied
	}
	if occupied > 4 {
		return Empty
	}

	return current
}

func (s *seatPlan) display() {

	fmt.Println()
	for i := 0; i < s.length; i++ {
		toPrint := ""
		for j := 0; j < s.width; j++ {
			toPrint += stateToValue(s.state(i, j))
		}
		fmt.Println(toPrint)
	}
	fmt.Println()
}

func valueToState(r rune) state {
	if r == '.' {
		return Floor
	}
	if r == 'L' {
		return Empty
	}
	if r == '#' {
		return Occupied
	}

	return Unset
}

func stateToValue(s state) string {
	if s == Floor {
		return "."
	}
	if s == Empty {
		return "L"
	}
	if s == Occupied {
		return "#"
	}
	return "X"
}

func fileAsSeatPlan() (s *seatPlan) {

	input := utils.FileAsLines("input/day11.txt")
	length := len(input)
	width := len(input[0])
	s = NewSeatPlan(width, length)

	for i, line := range input {
		for j, seat := range line {
			s.setState(j, i, valueToState(rune(seat)))
		}
	}

	return s
}

func countOccupied(seats *seatPlan) int {

	count := 0
	for i := 0; i < seats.width; i++ {
		for j := 0; j < seats.length; j++ {
			if seats.state(i, j) == Occupied {
				count += 1
			}
		}
	}
	return count
}

func nextSeatPlan(seats *seatPlan) (nextSeats *seatPlan, changed bool) {

	changed = false
	nextSeats = NewSeatPlan(seats.width, seats.length)

	for i := 0; i < seats.width; i++ {
		for j := 0; j < seats.length; j++ {
			nextSeats.setState(i, j, nextState(seats, i, j))

			if seats.state(i, j) != nextSeats.state(i, j) {
				changed = true
			}
		}
	}

	return
}

func SolveDay11() (string, string) {

	part1 := 0
	part2 := 0

	seats := fileAsSeatPlan()
	changed := true
	for changed {
		//seats.display()
		seats, changed = nextSeatPlan(seats)
	}

	part1 = countOccupied(seats)

	return utils.ResultStrings(part1, part2)
}
