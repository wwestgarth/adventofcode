package solutions

import (
	utils "adventofcode/utils"
	"strings"
)

type operation struct {
	name  string
	value int
}

type codeStack struct {
	ops      []*operation
	executed map[int]bool
	flipped  map[int]bool
	counter  int
}

func (c *codeStack) reset() {
	c.counter = 0
	c.executed = make(map[int]bool)
}

func (c *codeStack) flip(i int) {
	op := c.ops[i]

	if op.name == "jmp" {
		op.name = "nop"
	}
	if op.name == "nop" {
		op.name = "jmp"
	}
	return
}

func (c *codeStack) flipNext(i int) int {

	for index, op := range c.ops {

		if index <= i {
			continue
		}
		if op.name == "acc" {
			continue
		}

		c.flip(index)
		return index
	}

	return 0
}

func (c *codeStack) execute(line int) int {

	op := c.ops[line]

	if op.name == "jmp" {
		return line + op.value
	}

	if op.name == "acc" {
		c.counter += op.value
	}

	return line + 1
}

func (c *codeStack) run() bool {

	line := 0
	for {

		c.executed[line] = true
		line = c.execute(line)

		_, ok := c.executed[line]

		if ok {
			return false // ended old way
		}

		if line < 0 || line > len(c.ops) {
			return false
		}

		if line == len(c.ops) {
			return true // ended good way
		}
	}
}

func fileAsCode() codeStack {

	var code codeStack
	code.ops = []*operation{}
	code.executed = make(map[int]bool)
	code.flipped = make(map[int]bool)

	for _, input := range utils.FileAsLines("input/day8.txt") {

		op := new(operation)
		f := strings.Fields(input)

		op.name = f[0]
		op.value = utils.AtoiPanic(f[1])
		code.ops = append(code.ops, op)
	}

	return code

}

func SolveDay8() (string, string) {

	part1 := 0
	part2 := 0

	code := fileAsCode()
	code.run()
	part1 = code.counter

	code.reset()

	flipped := code.flipNext(0)
	for !code.run() {

		// flip it back
		code.flip(flipped)

		// reset counter things
		code.reset()

		// flip the next one along
		flipped = code.flipNext(flipped)
	}

	part2 = code.counter

	return utils.ResultStrings(part1, part2)
}
