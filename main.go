package main

import (
	"adventofcode/solutions"
	"flag"
	"fmt"
)

type options struct {
	day *int
}

func getOpts() options {

	var opts options

	// Choose day
	opts.day = flag.Int("day", 0, "which days solution to run")

	// Parse them
	flag.Parse()
	return opts

}

func main() {

	opts := getOpts()

	part1 := ""
	part2 := ""

	if *opts.day == 1 {
		part1, part2 = solutions.SolveDay1()
	}
	if *opts.day == 2 {
		part1, part2 = solutions.SolveDay2()
	}
	if *opts.day == 3 {
		part1, part2 = solutions.SolveDay3()
	}
	if *opts.day == 4 {
		part1, part2 = solutions.SolveDay4()
	}
	if *opts.day == 5 {
		part1, part2 = solutions.SolveDay5()
	}
	if *opts.day == 6 {
		part1, part2 = solutions.SolveDay6()
	}
	if *opts.day == 7 {
		part1, part2 = solutions.SolveDay7()
	}
	if *opts.day == 8 {
		part1, part2 = solutions.SolveDay8()
	}
	if *opts.day == 12 {
		part1, part2 = solutions.SolveDay12()
	}
	fmt.Println("Solutions")
	fmt.Println("Part1: ", part1)
	fmt.Println("Part2: ", part2)

}
