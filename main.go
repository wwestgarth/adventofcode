package main

import (
	day1 "adventofcode/solutions"
	day2 "adventofcode/solutions"
	day3 "adventofcode/solutions"
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
		part1, part2 = day1.SolveDay1()
	}
	if *opts.day == 2 {
		part1, part2 = day2.SolveDay2()
	}
	if *opts.day == 3 {
		part1, part2 = day3.SolveDay3()
	}
	fmt.Println("Solutions")
	fmt.Println("Part1: ", part1)
	fmt.Println("Part2: ", part2)

}
