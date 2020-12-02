package main

import (
	day1 "adventofcode/solutions"
	day2 "adventofcode/solutions"
	"flag"
	"fmt"
)

type options struct {
	day *int
}

func getOpts() options {

	var opts options

	// Functional choices
	opts.day = flag.Int("day", 0, "which days solution to run")

	// Parse them
	flag.Parse()
	return opts

}

func main() {

	opts := getOpts()

	res := ""
	if *opts.day == 1 {
		res = day1.SolveDay1()
	}
	if *opts.day == 2 {
		res = day2.SolveDay2()
	}
	fmt.Println("Solution:", res)

}
